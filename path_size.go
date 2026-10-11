package code

import (
	"code/internal/format"
	"code/internal/size"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := size.CalcPathSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	return format.FormatSize(size, human), nil
}
