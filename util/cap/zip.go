package cap

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func UnZipInMemory(zipPath string) (map[string][]byte, error) {
	zipData, err := os.ReadFile(zipPath)
	if err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, err
	}

	files := make(map[string][]byte)

	for _, zipFile := range zipReader.File {
		f, err := zipFile.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		println("unzip:", zipFile.Name)
		unzippedFileBytes, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		files[zipFile.Name] = unzippedFileBytes
	}

	return files, nil
}
