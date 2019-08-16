package assert_test

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, 0, 0)
	assert.NotEqual(t, 0, 1)
}
