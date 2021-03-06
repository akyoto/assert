package assert

import (
	"reflect"
	"testing"
)

// Nil asserts that the given parameter equals nil.
func Nil(t testing.TB, a interface{}) {
	if isNil(a) {
		return
	}

	t.Errorf(`assert.Nil:
value:    %v
expected: nil`, a)
	t.FailNow()
}

// NotNil asserts that the given parameter does not equal nil.
func NotNil(t testing.TB, a interface{}) {
	if !isNil(a) {
		return
	}

	t.Errorf(`assert.NotNil:
value:    %v
expected: not nil`, a)
	t.FailNow()
}

// isNil returns true if the object is nil.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)

	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return value.IsNil()
	}

	return false
}
