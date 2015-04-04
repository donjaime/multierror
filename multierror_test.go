package multierror

import (
	"fmt"
	"testing"
)

func TestZeroErrors(t *testing.T) {
	var e MultiError
	if e != nil {
		t.Error("should be nil")
	}
	if e.Error() != "0 errors" {
		t.Errorf("Expected '0 errors'. Got %s.", e.Error())
	}
}

func TestNilness(t *testing.T) {
	e := getNilErr()
	if e != nil {
		t.Error("should be nil")
	}
}

func getNilErr() error {
	var e MultiError
	return e.AsErr()
}

func TestNonZeroErrors(t *testing.T) {
	var mErr MultiError
	mErr = append(mErr, fmt.Errorf("An error"))
	if mErr == nil {
		t.Error("A nonempty MultiError should not be nil")
	}

	if len(mErr) != 1 {
		t.Error("The MultiError was not of length 1")
	}

	if mErr.Error() != "1 error: An error" {
		t.Error("MultiError (single) string was not as expected")
	}

	(&mErr).Add(fmt.Errorf("Another error"))
	if mErr.Error() != "2 errors: An error; Another error" {
		t.Error("MultiError (multiple) string was not as expected")
	}
}
