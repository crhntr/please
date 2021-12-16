package please

type testingT interface {
	Helper()
	Errorf(string, ...interface{})
	Logf(string, ...interface{})
}
