package multierror

import (
	"fmt"
	"testing"
)

func TestZeroErrors(t *testing.T) {
	var e MultiError
	if e != nil {
		t.Error("An empty Errors Err() method should return nil")
	}
}

func TestNonZeroErrors(t *testing.T) {
	var mErr MultiError
	mErr = append(mErr, fmt.Errorf("An error"))
	if mErr == nil {
		t.Error("A nonempty Errors Err() method should not return nil")
	}

	if len(mErr) != 1 {
		t.Error("The MultiError Errors field was not of length 1")
	}

	if mErr.Error() != "1 error: An error" {
		t.Error("MultiError (single) string was not as expected")
	}

	(&mErr).Add(fmt.Errorf("Another error"))
	if mErr.Error() != "2 errors: An error; Another error" {
		t.Error("MultiError (multiple) string was not as expected")
	}
}
