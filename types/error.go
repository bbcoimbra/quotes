package types

import (
	"fmt"
)

type NotFoundError struct {
	Id int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Quote with ID#%d not found", e.Id)
}
