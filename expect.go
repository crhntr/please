package please

func ExpectTrue(t testingT, got bool) bool {
	t.Helper()
	if !got {
		t.Errorf("\nexpectation failed, expected true, got:\n\t%t", got)
		return false
	}
	t.Logf("expectation met, got true")
	return true
}

func ExpectFalse(t testingT, got bool) bool {
	t.Helper()
	if got {
		t.Errorf("\nexpectation failed, expected false, got:\n\t%t", got)
		return false
	}
	t.Logf("expectation met, got false")
	return true
}
