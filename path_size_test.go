package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetPathSizeErrors(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"empty path", ""},
		{"nonexistent path", "testdata/does_not_exist"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetPathSize(test.path, false, false, false)

			require.Error(t, err)
			assert.Equal(t, "", got)
		})
	}
}

func TestHiddenRootIsMeasured(t *testing.T) {
	dir := t.TempDir()
	hidden := filepath.Join(dir, ".hidden")

	require.NoError(t, os.WriteFile(hidden, []byte("secret"), 0o644))

	size, err := GetPathSize(hidden, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "6B", size)
}

func TestHiddenChildren(t *testing.T) {
	dir := t.TempDir()

	require.NoError(t, os.WriteFile(filepath.Join(dir, "visible"), []byte("abc"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, ".hidden"), []byte("secret"), 0o644))

	size, err := GetPathSize(dir, false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "3B", size)

	sizeAll, err := GetPathSize(dir, false, false, true)
	require.NoError(t, err)
	assert.Equal(t, "9B", sizeAll)
}

func TestDotAsRoot(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "visible"), []byte("abc"), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, ".hidden"), []byte("secret"), 0o644))

	oldwd, err := os.Getwd()
	require.NoError(t, err)

	defer func() {
		require.NoError(t, os.Chdir(oldwd))
	}()

	require.NoError(t, os.Chdir(dir))

	size, err := GetPathSize(".", false, false, false)
	require.NoError(t, err)
	assert.Equal(t, "3B", size)
}
