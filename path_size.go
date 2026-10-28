// Package code - API подсчета и форматирования размера по указонному пути
package code

import (
	"github.com/Daaastal/go-project-242/internal/format"
	"github.com/Daaastal/go-project-242/internal/size"
)

// GetPathSize - получить размер по указанному пути
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := size.CalcPathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return format.Size(size, human), nil
}
