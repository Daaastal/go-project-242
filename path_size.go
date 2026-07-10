package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func convertHumanFormat(size int64) (string, error) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	i := 0

	for value >= 1024 && i < len(units)-1 {
		value /= 1024
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%d%s", size, units[i]), nil
	}

	return fmt.Sprintf("%.1f%s", value, units[i]), nil
}

func pathSize(path string, recursive, all bool) (int64, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if isHidden(fi.Name()) && !all {
		return 0, nil
	}

	if !fi.IsDir() {
		return fi.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, entry := range entries {
		if isHidden(entry.Name()) && !all {
			continue
		}
		entryPath := filepath.Join(path, entry.Name())
		var entryInfo os.FileInfo
		entryInfo, err = os.Lstat(entryPath)
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
			if recursive {
				var subSize int64
				subSize, err = pathSize(entryPath, recursive, all)
				if err != nil {
					return 0, err
				}
				size += subSize
			}
			continue
		}

		size += entryInfo.Size()
	}

	return size, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := pathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	if human {
		return convertHumanFormat(size)
	}

	return fmt.Sprintf("%dB", size), nil
}
