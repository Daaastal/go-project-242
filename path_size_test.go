package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	tests := []struct {
		expectErr  error
		expectSize string
		name       string
		path       string
		recursive  bool
		human      bool
		all        bool
	}{
		{
			nil, "0B",
			"empty file",
			"testdata/empty_file.txt",
			false, false, false,
		},
		{
			nil, "7B",
			"ordinary file",
			"testdata/hello_file.txt",
			false, false, false,
		},
		{
			nil, "0B",
			"empty dir",
			"testdata/empty_dir",
			false, false, false,
		},
		{
			nil, "20B",
			"dir without recursive",
			"testdata/dir",
			false, false, false,
		},
		{
			nil, "503B",
			"dir with recursive",
			"testdata/dir",
			true, false, false,
		},
		{
			nil, "89982B",
			"big file without human flag",
			"testdata/big_file.txt",
			false, false, false,
		},
		{
			nil, "87.9KB",
			"big file with human flag",
			"testdata/big_file.txt",
			false, true, false,
		},
		{
			nil, "14B",
			"symlink",
			"testdata/sym_hello_file",
			false, false, false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetPathSize(test.path,
				test.recursive,
				test.human,
				test.all)

			assert.Equal(t, test.expectSize, got)
			assert.Equal(t, test.expectErr, err)
		})
	}
}
