package domain

import (
	"fmt"
)

// A Validator represents a series of validation rules associated with a (type, version) combination.
// Such rules must follow a certain format defined by the current FormatChecker implementation
type Validator struct {
	Type    string
	Version int
	Rules   []byte
}

// NewValidator instantiates a new Validator for a certain type and with a certain set of rules. Such rules are
// checked upon creation and may result in an error. Otherwise, it returns a Validator instance with its version
// unassigned. It will remain unassigned until such validator is persisted by the repository and the next due
// version is determined.
func NewValidator(_type string, rules []byte) (*Validator, error) {
	err := Current.GetFormatChecker().Check(rules)
	if err != nil {
		return nil, ErrValidatorHasWrongFormat{_type, err.Error()}
	}

	return &Validator{
		Type:    _type,
		Version: VersionNotAssignedYet, // It will be set by the
		Rules:   rules,
	}, nil
}

// A FormatChecker is responsible for ensuring that a set of validation rules comply with a specific format
// (e.g. JSON Schema, XML Schema, or a serialization of a Golang function)
type FormatChecker interface {
	Check(rules []byte) error
}

// ErrValidatorHasWrongFormat is returned when trying to instantiate a validator whose format does not comply
// with the one defined by the current FormatChecker implementation
type ErrValidatorHasWrongFormat struct {
	Type, Cause string
}

func (e ErrValidatorHasWrongFormat) Error() string {
	return fmt.Sprintf("Validator for type '%s' has a wrong format: %s", e.Type, e.Cause)
}
