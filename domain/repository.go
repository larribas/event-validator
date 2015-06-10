package domain

import (
	"fmt"
)

type Repository interface {
	Create(validator *Validator) (version int)
	GetNextVersion(_type string) int
	Inspect(_type string, version int) (*Validator, error)
}

// ErrValidatorDoesNotExist is returned when the validator searched for is not present in the system
type ErrValidatorDoesNotExist struct {
	Type    string
	Version int
}

func (e ErrValidatorDoesNotExist) Error() string {
	return fmt.Sprintf("Validator for type '%s' with version %d does not exist", e.Type, e.Version)
}
