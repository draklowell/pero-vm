package internal

import (
	"fmt"
)

type ErrEntryNotFound struct{ Entry string }

func (e *ErrEntryNotFound) Error() string {
	return fmt.Sprintf("entry \"%s\" not found", e.Entry)
}
