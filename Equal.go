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

	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		t.Errorf(`assert.Equal:
value:    %v
expected: %v`, a, b)
	} else {
		t.Errorf(`assert.Equal:
value:    %v (%s)
expected: %v (%s)`, a, reflect.TypeOf(a).Name(), b, reflect.TypeOf(b).Name())
	}

	t.FailNow()
}

// NotEqual asserts that the two parameters are equal.
func NotEqual(t testing.TB, a interface{}, b interface{}) {
	if a != b {
		return
	}

	t.Errorf(`assert.NotEqual:
value:    %v
expected: %v`, a, b)
	t.FailNow()
}

// DeepEqual asserts that the two parameters are deeply equal.
func DeepEqual(t testing.TB, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf(`assert.DeepEqual:
value:    %v
expected: %v`, a, b)
	t.FailNow()
}
