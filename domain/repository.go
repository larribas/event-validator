package domain

import (
	"fmt"
)

// VersionNotAssignedYet identifies the next version for an empty type (one with no validators)
const VersionNotAssignedYet = -1

// A Repository is responsible for the storage and retrieval of validators from a database (or other persistence mechansim)
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
