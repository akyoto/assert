package assert

import "testing"

// True asserts that the given parameter is true.
func True(t testing.TB, a interface{}) {
	Equal(t, a, true)
}

// False asserts that the given parameter is false.
func False(t testing.TB, a interface{}) {
	Equal(t, a, false)
}
