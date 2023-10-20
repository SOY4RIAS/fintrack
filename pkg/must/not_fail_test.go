package must_test

import (
	"errors"
	"fintrack/pkg/must"
	"testing"
)

func TestCallsFunctionWithoutError(_ *testing.T) {
	must.NotFailFn(func() error {
		return nil
	})
}

// When given a nil error, it should not panic.
func TestNotFailWithNilError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected no panic, but got: %v", r)
		}
	}()
	must.NotFail(nil)
}

// When given a non-nil error, it should panic with the given error.
func TestNotFailWithNonNilError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but got none")
		} else if err, ok := r.(error); !ok || err.Error() != "test error" {
			t.Errorf("Expected panic with error 'test error', but got: %v", r)
		}
	}()
	must.NotFail(errors.New("test error"))
}
