//go:build go1.18

package please

func ExpectNilError(t testingT, err error) bool {
	t.Helper()
	if err != nil {
		t.Errorf("expectation failed, got an error: %s", err)
		return false
	}
	t.Logf("expectation met, error is nil")
	return true
}

func ExpectEqual[T comparable](t testingT, got, exp T) bool {
	t.Helper()
	if got != exp {
		t.Errorf("\nexpectation failed, got:\n\t%v\nbut expected:\n\t%v", got, exp)
		return false
	}
	t.Logf("expectation met, %v is equal to %v", exp, got)
	return true
}
