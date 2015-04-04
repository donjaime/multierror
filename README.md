# multierror #

multierror is a simple Go package for combining multiple `error`s.
This is handy if you are concurrently running operations within
a function that returns only a single `error`.

## API ##

multierror exposes just one type.

`multierror.MultiError` implements the `error` interface.  It is
also an alias of `[]error`. So `multierror.MultiError` can be appended to.

`*multierror.MultiError` also exposes an `Add(error)` method which is slightly
 nicer than using the `mErr = append(mErr, err)` idiom.

## Example ##

```go
package main

import (
	"fmt"
	"github.com/donjaime/multierror"
)

func main() {
	// Collect multiple errors together in multierror.Errors
	var errs multierror.MultiError
	errs = append(e1, fmt.Errorf("Error 1"))
	(&errs).Add(fmt.Errorf("Error 2"))

	// Output: "2 errors: Error 1; Error 2"
	fmt.Println(errs)

	// Iterate over the individual errors
	for _, err := range errs {
		fmt.Println(err) // Output: "Error 1" and "Error 2"
	}
}
```
