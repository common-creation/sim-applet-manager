package gp

import (
	"bufio"
	"context"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/common-creation/sim-applet-manager/util/apdu"
	"github.com/common-creation/sim-applet-manager/util/apppath"
	"github.com/common-creation/sim-applet-manager/util/cap"
	"github.com/common-creation/sim-applet-manager/util/command"
	"github.com/common-creation/sim-applet-manager/util/db"
	"github.com/common-creation/sim-applet-manager/util/i18n"
	"github.com/common-creation/sim-applet-manager/util/log"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	RunOption struct {
		CardReader string   `json:"cardReader"`
		GpArgs     []string `json:"args"`
	}
	Result struct {
		Success bool   `json:"success"`
		Output  string `json:"output"`
		Status  string `json:"status"`
	}
	ListResult struct {
		Package HexFingerPrint   `json:"package"`
		Applets []HexFingerPrint `json:"applets"`
	}
	HexFingerPrint struct {
		Hex         string `json:"hex"`
		FingerPrint string `json:"fingerPrint"`
	}
)

func Path(ctx context.Context, i18n *i18n.I18n) string {
	gpPath := ""

	if runtime.GOOS == "darwin" {
		if stdout, err := exec.Command("command", "-v", "gp").Output(); err == nil {
			gpPath = strings.TrimSpace(string(stdout))
		}
		if _, err := exec.Command(gpPath, "--help").Output(); err == nil {
			return gpPath
		}
		if stdout, err := exec.Command("brew", "--prefix").Output(); err == nil {
			gpPath = filepath.Join(strings.TrimSpace(string(stdout)), "bin", "gp")
		}
		if _, err := exec.Command(gpPath, "--help").Output(); err == nil {
			return gpPath
		}
		if stdout, err := exec.Command("/opt/homebrew/bin/brew", "--prefix").Output(); err == nil {
			gpPath = filepath.Join(strings.TrimSpace(string(stdout)), "bin", "gp")
		}
		if _, err := exec.Command(gpPath, "--help").Output(); err == nil {
			return gpPath
		}
		if stdout, err := exec.Command("/usr/local/bin/brew", "--prefix").Output(); err == nil {
			gpPath = filepath.Join(strings.TrimSpace(string(stdout)), "bin", "gp")
		}
		if _, err := exec.Command(gpPath, "--help").Output(); err == nil {
			return gpPath
		}
	} else {
		cmd1 := exec.Command("gp.exe", "--help")
		command.HideWindow(cmd1)

		gpPath = filepath.Join(apppath.MustAppDirPath(ctx, i18n), "gp.exe")
		cmd2 := exec.Command(gpPath, "--help")
		command.HideWindow(cmd2)
		if _, err := cmd2.Output(); err == nil {
			return gpPath
		}
	}

	return ""
}

func Run(ctx context.Context, i18n *i18n.I18n, option RunOption) Result {
	gpPath := Path(ctx, i18n)
	if gpPath == "" {
		// TODO: エラーダイアログ
		return Result{
			Success: false,
		}
	}

	args := []string{"-dv", "-r", option.CardReader}
	args = append(args, option.GpArgs...)

	commandCtx := context.TODO()
	cmd := exec.CommandContext(commandCtx, gpPath, args...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		// TODO: エラーダイアログ
		panic(err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		// TODO: エラーダイアログ
		panic(err)
	}

	command.HideWindow(cmd)

	if err := cmd.Start(); err != nil {
		// TODO: エラーダイアログ
		panic(err)
	}

	outputBuilder := strings.Builder{}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			line := scanner.Text()
			println(line)
			wailsRuntime.EventsEmit(ctx, "gpLogs", line)
			outputBuilder.WriteString(line + "\n")
		}
		if err := scanner.Err(); err != nil {
			// TODO: エラーダイアログ
			println(err)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			line := scanner.Text()
			println(line)
			wailsRuntime.EventsEmit(ctx, "gpLogs", line)
			outputBuilder.WriteString(line + "\n")
		}
		if err := scanner.Err(); err != nil {
			// TODO: エラーダイアログ
			println(err)
		}
	}()

	if err := cmd.Wait(); err != nil {
		// TODO: エラーダイアログ
		println(err)
	}
	println("done!")

	log.WriteString(ctx, i18n, outputBuilder.String())

	return Result{
		Success: cmd.ProcessState.Success(),
		Output:  outputBuilder.String(),
	}
}

func GetICCID(ctx context.Context, i18n *i18n.I18n, cardReader string) string {
	result := Run(ctx, i18n, RunOption{
		CardReader: cardReader,
		GpArgs: []string{
			"-a", "A0A40000023F00",
			"-a", "A0A40000022FE2",
			"-a", "A0B000000A",
		},
	})

	found := false
	for _, line := range strings.Split(result.Output, "\n") {
		trimmed := strings.ReplaceAll(line, " ", "")
		println("line:", line)
		println("trimmed:", trimmed)
		if strings.HasPrefix(trimmed, "A>>") && strings.HasSuffix(trimmed, "A0B000000A") {
			found = true
			continue
		}
		if found {
			elems := strings.Split(line, " ")
			println(elems)
			if elems[len(elems)-1] == "9000" {
				origIccid := elems[len(elems)-2]
				swappedIccid := apdu.SwapPairs(origIccid)
				return strings.TrimSuffix(swappedIccid, "F")
			}
		}
	}
	return ""
}

func extractPackageAidHex(input string) string {
	re := regexp.MustCompile(`[A-F0-9]{32}`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return ""
	}
	return matches[0]
}

func extractPackageFingerPrint(input string) string {
	re := regexp.MustCompile(`\|([^\|]+)\|`)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		return ""
	}
	return strings.Trim(matches[0], "|")
}

func ListApplets(ctx context.Context, i18n *i18n.I18n, cardReader string, key db.Key) []ListResult {
	result := Run(ctx, i18n, RunOption{
		CardReader: cardReader,
		GpArgs: []string{
			"--connect", key.AID,
			"--key-enc", key.EncKey,
			"--key-mac", key.MacKey,
			"--key-dek", key.KekKey,
			"--list",
		},
	})

	applets := make([]ListResult, 0)
	var pkg *ListResult
	for _, line := range strings.Split(result.Output, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "PKG:") {
			if pkg != nil {
				applets = append(applets, *pkg)
			}

			pkg = &ListResult{}
			pkg.Applets = make([]HexFingerPrint, 0)
			pkg.Package = HexFingerPrint{
				Hex:         extractPackageAidHex(trimmed),
				FingerPrint: extractPackageFingerPrint(trimmed),
			}
			continue
		}
		if strings.HasPrefix(trimmed, "Applet:") {
			pkg.Applets = append(pkg.Applets, HexFingerPrint{
				Hex:         extractPackageAidHex(trimmed),
				FingerPrint: extractPackageFingerPrint(trimmed),
			})
			continue
		}
	}
	if pkg != nil {
		applets = append(applets, *pkg)
	}
	return applets
}

func UninstallApplet(ctx context.Context, i18n *i18n.I18n, cardReader string, key db.Key, aid string) Result {
	return Run(ctx, i18n, RunOption{
		CardReader: cardReader,
		GpArgs: []string{
			"--connect", key.AID,
			"--key-enc", key.EncKey,
			"--key-mac", key.MacKey,
			"--key-dek", key.KekKey,
			"--delete", aid,
			"--force",
		},
	})
}

func InstallApplet(ctx context.Context, i18n *i18n.I18n, cardReader string, key db.Key, capPath string, params string) Result {
	capFiles, err := cap.UnZipInMemory(capPath)
	if err != nil {
		return Result{
			Success: false,
			Output:  err.Error(),
		}
	}
	manifest := cap.ManifestToMap(string(capFiles["META-INF/MANIFEST.MF"]))
	aid := extractPackageAidHex(strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(manifest["Java-Card-Package-AID"], "0x", ""), ":", "")))
	println("AID:", aid)

	pkgs := ListApplets(ctx, i18n, cardReader, key)

	for _, pkg := range pkgs {
		println(pkg.Package.Hex, aid, pkg.Package.Hex == aid)
		if pkg.Package.Hex == aid {
			UninstallApplet(ctx, i18n, cardReader, key, pkg.Package.Hex)
		}
	}

	gpArgs := []string{
		"--connect", key.AID,
		"--key-enc", key.EncKey,
		"--key-mac", key.MacKey,
		"--key-dek", key.KekKey,
		"--install", capPath,
	}
	if params != "" {
		gpArgs = append(gpArgs, "--param", params)
	}
	return Run(ctx, i18n, RunOption{
		CardReader: cardReader,
		GpArgs:     gpArgs,
	})
}
