// Package code - API подсчета и форматирования размера по указонному пути
package code

import (
	"code/internal/format"
	"code/internal/size"
)

// GetPathSize - получить размер по указанному пути
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := size.CalcPathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return format.HumanSize(size, human), nil
}
