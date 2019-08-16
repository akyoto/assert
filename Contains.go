package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Contains asserts that a contains b.
func Contains(t *testing.T, a interface{}, b interface{}) {
	if contains(a, b) {
		return
	}

	t.Errorf("assert.Contains:\nvalue: %v\ncollection: %v", a, b)
	t.FailNow()
}

// NotContains asserts that a doesn't contain b.
func NotContains(t *testing.T, a interface{}, b interface{}) {
	if !contains(a, b) {
		return
	}

	t.Errorf("assert.NotContains:\nvalue: %v\ncollection: %v", a, b)
	t.FailNow()
}

// contains returns whether container contains the given the element.
// It works with strings, maps and slices.
func contains(container interface{}, element interface{}) bool {
	value := reflect.ValueOf(container)

	switch value.Kind() {
	case reflect.String:
		elementValue := reflect.ValueOf(element)
		return strings.Contains(value.String(), elementValue.String())

	case reflect.Map:
		keys := value.MapKeys()

		for _, key := range keys {
			if key.Interface() == element {
				return true
			}
		}

	default:
		for i := 0; i < value.Len(); i++ {
			if value.Index(i).Interface() == element {
				return true
			}
		}
	}

	return false
}
