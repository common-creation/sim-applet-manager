package cardreader

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/samber/lo"
)

func ListCardReader() []string {
	if runtime.GOOS == "darwin" {
		return listCardReaderMacOS()
	} else {
		return listCardReaderWindows()
	}
}

func listCardReaderMacOS() []string {
	cardReaders := make([]string, 0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))
	defer cancel()
	stdoutReader := &strings.Builder{}

	cmd := exec.CommandContext(ctx, "pcsctest")
	cmd.Stdout = stdoutReader
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		println(err)
		return cardReaders
	}

	lines := strings.Split(stdoutReader.String(), "\n")
	readers := lo.FilterMap[string](lines, func(item string, _ int) (string, bool) {
		if strings.HasPrefix(item, "Reader ") {
			return strings.TrimSpace(strings.Split(item, ":")[1]), true
		}
		return "", false
	})
	cardReaders = append(cardReaders, readers...)
	sort.Strings(cardReaders)

	return cardReaders
}

var mutex = sync.Mutex{}

func listCardReaderWindows() []string {
	cardReaders := make([]string, 0)

	mutex.Lock()
	defer mutex.Unlock()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitialize(0); err != nil {
		println(err)
		return cardReaders
	}
	defer ole.CoUninitialize()

	// WMIのCOMオブジェクトを取得
	unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
	if err != nil {
		println(err)
	}
	defer unknown.Release()

	// IUnknownをIDispatchにクエリ
	wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		println(err)
	}
	defer wmi.Release()

	// WMIサービスに接続
	serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer")
	if err != nil {
		println(err)
	}
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	query := "SELECT * FROM Win32_PnPEntity WHERE DeviceID LIKE 'SWD\\\\SCDEVICEENUM\\\\%'"
	resultRaw, err := oleutil.CallMethod(service, "ExecQuery", query)
	if err != nil {
		println(err)
	}
	result := resultRaw.ToIDispatch()
	defer result.Release()

	enumProperty, err := oleutil.GetProperty(result, "_NewEnum")
	if err != nil {
		println(err)
	}
	enum, err := enumProperty.ToIUnknown().IEnumVARIANT(ole.IID_IEnumVariant)
	if err != nil {
		println(err)
	}
	defer enum.Release()

	for {
		itemRaw, fetched, err := enum.Next(1)
		if err != nil || fetched == 0 {
			break
		}
		item := itemRaw.ToIDispatch()
		defer item.Release()

		caption, _ := oleutil.GetProperty(item, "Caption")
		cardReaders = append(cardReaders, caption.Value().(string))
	}

	return cardReaders
}

func ResetOnMacOS(cardReader string) {
	if runtime.GOOS != "darwin" {
		return
	}

	// NOTE: gpでICCIDを読み取るとなぜか5秒くらい次の読み取りができなくなるけどpcsctestを叩くとうまくいく
	cmd := exec.Command("pcsctest")
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		println(err)
		return
	}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		println(err)
		return
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		println(err)
		return
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		readers := 0
		readerIndex := -1
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			println(line)
			if strings.Contains(line, "Waiting for card insertion") {
				_ = cmd.Process.Kill()
				break
			}
			if strings.HasPrefix(line, "Reader ") {
				readers++
				readerName := strings.TrimSpace(strings.Split(line, ":")[1])
				if readerName == cardReader {
					readerIndex = readers
					stdinPipe.Write([]byte(strconv.Itoa(readerIndex) + "\n"))
				}
			}
			if strings.HasPrefix(line, "Enter the reader number ") {
				stdinPipe.Write([]byte(strconv.Itoa(readerIndex) + "\n"))
			}
		}
		if err := scanner.Err(); err != nil {
			// TODO: エラーダイアログ
			println(err)
		}
	}()

	_ = cmd.Wait()
}
