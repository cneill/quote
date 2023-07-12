package prettify_test

import (
	"testing"

	"github.com/cneill/quote/pkg/prettify"

	"github.com/stretchr/testify/assert"
)

func TestReplaceGarbage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{"This is an examp-\nle of detec-\nting wrapping", "This is an example of detecting wrapping"},
		{"Stupid ”garbage quotes“", "Stupid \"garbage quotes\""},
		{"Stupid fancy dash —  and  — double space", "Stupid fancy dash - and - double space"},
		{"I\ndon’t\nwant\nnewlines", "I don't want newlines"},
	}

	for _, test := range tests {
		actual := prettify.ReplaceGarbage(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
