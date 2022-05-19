//go:build typeparams

package please

func ExpectEqual[T comparable](t testingT, got, exp T) bool {
	t.Helper()
	if got != exp {
		t.Errorf("\nexpectation failed, got:\n\t%v\nbut expected:\n\t%v", got, exp)
		return false
	}
	t.Logf("expectation met, %v is equal to %v", exp, got)
	return true
}
