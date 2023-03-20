package tests

import "testing"

func assertError(t testing.TB, name string, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got error %q want %q", name, got, want)
	}
}

func assertFloat64(t testing.TB, name string, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g want %g", name, got, want)
	}
}
