//go:build !typeparams

package please

import "reflect"

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

	gotStr, gotStrOK := got.(string)
	expStr, expStrOK := exp.(string)
	if gotStrOK && expStrOK {
		if gotStr != expStr {
			t.Errorf("\nexpectation failed, got:\n\t%q\nbut expected:\n\t%q", got, exp)
			return false
		}
		t.Logf("expectation met, got %q", got)
		return true
	}

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\nexpectation failed, got:\n\t%v\nbut expected:\n\t%v", got, exp)
		return false
	}
	t.Logf("expectation met, got %v", got)
	return true
}
