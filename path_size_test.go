package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetPathSize(t *testing.T) {
	tests := []struct {
		name		string
		path		string
		recursive	bool
		human		bool
		all		bool
		expect_size	string
		expect_err	error
	}{
		{
			"empty file",
			"testdata/empty_file.txt",
			false, false, false,
			"0B", nil,
		},
		{
			"ordinary file",
			"testdata/hello_file.txt",
			false, false, false,
			"7B", nil,
		},
		{
			"empty dir",
			"testdata/empty_dir",
			false, false, false,
			"0B", nil,
		},
		{
			"dir without recursive",
			"testdata/dir",
			false, false, false,
			"20B", nil,
		},
		{
			"dir with recursive",
			"testdata/dir",
			true, false, false,
			"503B", nil,
		},
		{
			"big file without human flag",
			"testdata/big_file.txt",
			false, false, false,
			"89982B", nil,
		},
		{
			"big file witht human flag",
			"testdata/big_file.txt",
			false, true, false,
			"87.9KB", nil,
		},
		{
			"big file witht human flag",
			"testdata/big_file.txt",
			false, true, false,
			"87.9KB", nil,
		},
		{
			"symlink",
			"testdata/sym_hello_file",
			false, false, false,
			"14B", nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetPathSize(test.path,
						test.recursive,
						test.human,
						test.all)

			assert.Equal(t, test.expect_size, got)
			assert.Equal(t, test.expect_err, err)
		})
	}
}
