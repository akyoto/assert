package assert

import (
	"testing"
)

// Nil asserts that the given parameter equals nil.
func Nil(t *testing.T, a interface{}) {
	if a == nil {
		return
	}

	t.Errorf("assert.Nil:\nvalue: %v\nexpected nil value", a)
	t.FailNow()
}

// NotNil asserts that the given parameter does not equal nil.
func NotNil(t *testing.T, a interface{}) {
	if a != nil {
		return
	}

	t.Errorf("assert.NotNil:\nvalue: %v\nexpected: non-nil value", a)
	t.FailNow()
}
