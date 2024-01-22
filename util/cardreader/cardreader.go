package cardreader

import (
	"context"
	"os"
	"os/exec"
	"runtime"
	"sort"
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
