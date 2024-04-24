package mongoquery

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singular(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"bodies":  "body",
		"classes": "class",
		"tests":   "test",
		"single":  "single",
		"":        "",
	}

	for plural, expected := range tests {
		result := singular(plural)

		require.Equal(t, expected, result)
	}
}
