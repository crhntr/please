package please

type testingT interface {
	Helper()
	Errorf(string, ...any)
	Logf(string, ...any)
}
