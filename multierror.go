// multierror is a simple Go package for combining multiple errors together.
package multierror

import (
	"bytes"
	"fmt"
)

// Type alias for a slice of errors
type MultiError []error

// Satisfy the error built-in interface
func (m MultiError) Error() string {
	if len(m) == 0 {
		return "0 errors"
	}

	var buf bytes.Buffer

	if len(m) == 1 {
		buf.WriteString("1 error: ")
	} else {
		fmt.Fprintf(&buf, "%d errors: ", len(m))
	}

	for i, err := range m {
		if i != 0 {
			buf.WriteString("; ")
		}

		buf.WriteString(err.Error())
	}

	return buf.String()
}

// To work around Go's nasty nil becomes not-nil when passed as interface across function boundary.
func (m MultiError) AsErr() error {
	if len(m) == 0 {
		return nil
	}
	return m
}

// Sugar for adding an error. Method only valid for pointers to MultiError.
func (m *MultiError) Add(err error) {
	*m = append(*m, err)
}
