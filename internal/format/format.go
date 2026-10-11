package format

import (
	"fmt"
)

func GetStringResult(size, path string) string {
	return fmt.Sprintf("%s\t%s\n", size, path)
}

func FormatSize(size int64, human bool) string {
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

	if i == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[i])
}
