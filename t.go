package please

type testingT interface {
	Helper()
	Errorf(arg1 string, arg2 ...interface{})
	Logf(arg1 string, arg2 ...interface{})
}
