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

func ExpectNilError(t testingT, err error) bool {
	t.Helper()
	if err != nil {
		t.Errorf("expectation failed, got an error: %s", err)
		return false
	}
	t.Logf("expectation met, error is nil")
	return true
}
