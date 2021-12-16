//go:build !typeparams

package please

import "reflect"

type any = interface{}

func ExpectNilError(t testingT, err error) bool {
	t.Helper()
	if err != nil {
		t.Errorf("expectation failed, got an error: %s", err)
		return false
	}
	t.Logf("expectation met, error is nil")
	return true
}

func ExpectEqual(t testingT, got, exp interface{}) bool {
	t.Helper()
	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\nexpectation failed, got:\n\t%v\nbut expected:\n\t%v", got, exp)
		return false
	}
	t.Logf("expectation met, %v is equal to %v", exp, got)
	return true
}
