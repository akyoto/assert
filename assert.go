package assert

import "testing"

// Equal asserts that the two parameters are equal.
func Equal(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}

	t.Errorf("assert.Equal:\nvalue: %v\nexpected: %v", a, b)
	t.FailNow()
}

// NotEqual asserts that the two parameters are equal.
func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		return
	}

	t.Errorf("assert.NotEqual:\nvalue: %v\nexpected: %v", a, b)
	t.FailNow()
}
