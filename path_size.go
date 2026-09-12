package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func formatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	i := 0

	if !human {
		return fmt.Sprintf("%dB", size)
	}

	for value >= 1024 && i < len(units)-1 {
		value /= 1024
		i++
	}

	return fmt.Sprintf("%.1f%s", value, units[i])
}

func pathSize(path string, recursive, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("lstat %q: %w", info, err)
	}

	mode := info.Mode()

	if mode.IsDir() {
		return dirSize(path, recursive, all)
	}

	if mode.IsRegular() || (mode&os.ModeType == os.ModeSymlink) {
		return info.Size(), nil
	}

	return 0, nil
}

func dirSize(path string, recursive, all bool) (int64, error) {
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

		subSize, err := pathSize(filepath.Join(path, name), recursive, all)
		if err != nil {
			return 0, err
		}
		size += subSize
	}
	return size, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := pathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}
