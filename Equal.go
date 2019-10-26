package assert

import (
	"reflect"
	"testing"
)

// Equal asserts that the two parameters are equal.
func Equal(t testing.TB, a interface{}, b interface{}) {
	if a == b {
		return
	}

	t.Errorf("assert.Equal:\nvalue: %v\nexpected: %v", a, b)
	t.FailNow()
}

// NotEqual asserts that the two parameters are equal.
func NotEqual(t testing.TB, a interface{}, b interface{}) {
	if a != b {
		return
	}

	t.Errorf("assert.NotEqual:\nvalue: %v\nexpected: %v", a, b)
	t.FailNow()
}

// DeepEqual asserts that the two parameters are deeply equal.
func DeepEqual(t testing.TB, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf("assert.DeepEqual:\nvalue: %v\nexpected: %v", a, b)
	t.FailNow()
}
