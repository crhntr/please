package please

import (
	"errors"
	"testing"

	"github.com/crhntr/please/fakes"
)

func TestExpectEqual(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tt := new(fakes.TestingT)

		result := ExpectEqual(tt, 10, 10)
		isTrue(t, result)
		isHelper(t, tt)
		isNotError(t, tt)
		isLogCalled(t, tt)
	})

	t.Run("failed", func(t *testing.T) {
		tt := new(fakes.TestingT)

		result := ExpectEqual(tt, 7, 10)
		isFalse(t, result)
		isHelper(t, tt)
		isError(t, tt)
		isLogNotCalled(t, tt)
	})
}

func TestExpectNilError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tt := new(fakes.TestingT)

		var err error
		result := ExpectNilError(tt, err)
		isTrue(t, result)
		isHelper(t, tt)
		isNotError(t, tt)
		isLogCalled(t, tt)
	})

	t.Run("failure", func(t *testing.T) {
		tt := new(fakes.TestingT)

		err := errors.New("banana")
		result := ExpectNilError(tt, err)
		isFalse(t, result)
		isHelper(t, tt)
		isError(t, tt)
		isLogNotCalled(t, tt)
	})
}

func isHelper(t *testing.T, tt *fakes.TestingT) {
	t.Helper()

	if tt.HelperCallCount() != 1 {
		t.Error("expected Helper to be called")
	}
}

func isError(t *testing.T, tt *fakes.TestingT) {
	t.Helper()

	if tt.ErrorfCallCount() != 1 {
		t.Error("expected expectation to fail")
	}
}

func isNotError(t *testing.T, tt *fakes.TestingT) {
	t.Helper()

	if tt.ErrorfCallCount() != 0 {
		t.Error("expected expectation to be met")
	}
}

func isLogCalled(t *testing.T, tt *fakes.TestingT) {
	t.Helper()

	if tt.LogfCallCount() != 1 {
		t.Error("expected some log")
	}
}

func isLogNotCalled(t *testing.T, tt *fakes.TestingT) {
	t.Helper()

	if tt.LogfCallCount() != 0 {
		t.Errorf("expected no logging, got %d calls", tt.LogfCallCount())
	}
}

func isTrue(t *testing.T, b bool) {
	t.Helper()

	if !b {
		t.Error("expected true, got false")
	}
}

func isFalse(t *testing.T, b bool) {
	t.Helper()

	if b {
		t.Error("expected false, got true")
	}
}
