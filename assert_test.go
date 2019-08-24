package assert_test

import (
	"testing"

	"github.com/akyoto/assert"
)

type testType struct {
	A int
}

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
	assert.NotNil(t, 0)
	assert.NotNil(t, "Hello")
	assert.NotNil(t, testType{})
	assert.NotNil(t, &testType{})
}

func TestEqual(t *testing.T) {
	assert.Equal(t, 0, 0)
	assert.NotEqual(t, 0, 1)
	assert.Equal(t, "Hello", "Hello")
	assert.NotEqual(t, "Hello", "World")
	assert.NotEqual(t, &testType{A: 10}, &testType{A: 10})
	assert.NotEqual(t, testType{A: 10}, testType{A: 20})
	assert.DeepEqual(t, testType{A: 10}, testType{A: 10})
}

func TestContains(t *testing.T) {
	assert.Contains(t, "Hello", "H")
	assert.Contains(t, "Hello", "Hello")
	assert.NotContains(t, "Hello", "404")
	assert.Contains(t, []string{"Hello", "World"}, "Hello")
	assert.NotContains(t, []string{"Hello", "World"}, "404")
	assert.Contains(t, []int{1, 2, 3}, 2)
	assert.NotContains(t, []int{1, 2, 3}, 4)
	assert.Contains(t, []int{1, 2, 3}, []int{})
	assert.Contains(t, []int{1, 2, 3}, []int{1, 2})
	assert.NotContains(t, []int{1, 2, 3}, []int{2, 1})
	assert.Contains(t, []byte{'H', 'e', 'l', 'l', 'o'}, byte('e'))
	assert.NotContains(t, []byte{'H', 'e', 'l', 'l', 'o'}, byte('a'))
	assert.Contains(t, []byte{'H', 'e', 'l', 'l', 'o'}, []byte{'e', 'l'})
	assert.NotContains(t, []byte{'H', 'e', 'l', 'l', 'o'}, []byte{'l', 'e'})
	assert.Contains(t, map[string]int{"Hello": 1, "World": 2}, "Hello")
	assert.NotContains(t, map[string]int{"Hello": 1, "World": 2}, "404")
}
