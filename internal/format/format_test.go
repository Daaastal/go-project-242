package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringResult(t *testing.T) {
	expected := "one\ttwo\n"
	actual := GetStringResult("one", "two")
	assert.Equal(t, expected, actual)
}

func TestHumanSize(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		size     int64
		human    bool
	}{
		{
			name:     "BigSizeWithNoHuman",
			expected: "89982B",
			size:     89982,
			human:    false,
		},
		{
			name:     "BigSizeWithHuman",
			expected: "87.9KB",
			size:     89982,
			human:    true,
		},
		{
			name:     "SmallSizeWithNoHuman",
			expected: "10B",
			size:     10,
			human:    false,
		},
		{
			name:     "SmallSizeWithHuman",
			expected: "10B",
			size:     10,
			human:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := HumanSize(test.size, test.human)
			assert.Equal(t, test.expected, got)
		})
	}
}
