package formaterror

import (
	"errors"
)

// FormatError is...
func FormatError(err string) error {
	return errors.New(err)
}
