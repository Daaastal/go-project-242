// Package format - API форматирования результата и размера
package format

import (
	"fmt"
	"math"
)

// GetStringResult - получить форматированный результат
func GetStringResult(size, path string) string {
	return fmt.Sprintf("%s\t%s\n", size, path)
}

// HumanSize - отформатировать размер с учетом флага human
func HumanSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)

	if !human {
		return fmt.Sprintf("%dB", size)
	}

	var selectedUnit string
	for _, unit := range units {
		rounded := math.Round(value*10) / 10
		if rounded < 1024 || unit == units[len(units)-1] {
			selectedUnit = unit

			break
		}
		value /= 1024
	}

	if selectedUnit == "B" {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, selectedUnit)
}
