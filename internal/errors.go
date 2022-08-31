package internal

import (
	"fmt"
)

type ErrorEntryNotFound struct{ Entry string }

func (e *ErrorEntryNotFound) Error() string {
	return fmt.Sprintf("EntryNotFound: %s", e.Entry)
}
