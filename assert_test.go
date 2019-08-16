package assert_test

import (
	"testing"

	"github.com/akyoto/assert"
)

type testType struct {
	A int
}

func TestEqual(t *testing.T) {
	assert.Equal(t, 0, 0)
	assert.NotEqual(t, 0, 1)
	assert.Equal(t, "Hello", "Hello")
	assert.NotEqual(t, "Hello", "World")
	assert.NotEqual(t, &testType{A: 10}, &testType{A: 20})
	assert.NotEqual(t, testType{A: 10}, testType{A: 20})
	assert.DeepEqual(t, testType{A: 10}, testType{A: 10})
}

func TestContains(t *testing.T) {
	assert.Contains(t, "Hello", "H")
	assert.Contains(t, "Hello", "Hello")
	assert.NotContains(t, "Hello", "404")
	assert.Contains(t, []string{"Hello", "World"}, "Hello")
	assert.NotContains(t, []string{"Hello", "World"}, "404")
	assert.Contains(t, map[string]int{"Hello": 1, "World": 2}, "Hello")
	assert.NotContains(t, map[string]int{"Hello": 1, "World": 2}, "404")
}

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
	assert.NotNil(t, 0)
	assert.NotNil(t, "Hello")
	assert.NotNil(t, testType{})
	assert.NotNil(t, &testType{})
}
