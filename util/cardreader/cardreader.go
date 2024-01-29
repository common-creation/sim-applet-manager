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
	"time"

	"github.com/samber/lo"
)

func ListCardReader() []string {
	cardReaders := make([]string, 0)

	if runtime.GOOS == "darwin" {
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
	} else {
		// TODO: Windows
		return cardReaders
	}
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
