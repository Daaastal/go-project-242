// Package size - API для получения размера
package size

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ErrUnsupportedFileType - тип файла не поддерживается для измерения
var ErrUnsupportedFileType = errors.New("unsupported file type")

// CalcPathSize - получить размер с учетом флагов
func CalcPathSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	mode := info.Mode()

	if mode.IsDir() {
		return calcDirSize(path, recursive, all)
	}

	if mode.IsRegular() || (mode&os.ModeType != 0) {
		return info.Size(), nil
	}

	return 0, fmt.Errorf("%w: %s", ErrUnsupportedFileType, path)
}

func calcDirSize(path string, recursive, all bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("read dir %q: %w", path, err)
	}

	var size int64
	for _, entry := range entries {
		name := entry.Name()

		if !all && strings.HasPrefix(name, ".") {
			continue
		}

		if !recursive && entry.IsDir() {
			continue
		}

		subSize, errSub := CalcPathSize(filepath.Join(path, name), recursive, all)
		if errSub != nil {
			return 0, errSub
		}
		size += subSize
	}

	return size, nil
}
