package domain

import (
    "fmt"
)

const VersionNotAssignedYet = -1

type Validator struct {
    Type string
    Version int
    Rules []byte
}

func NewValidator(_type string, rules []byte) (*Validator, error) {
    err := Current.GetFormatChecker().Check(rules)
    if err != nil {
        return nil, ErrValidatorHasWrongFormat{_type, err.Error()}
    }

    return &Validator{
        Type: _type,
        Version: VersionNotAssignedYet, // It will be set by the
        Rules: rules,
    }, nil
}

type FormatChecker interface {
    Check(rules []byte) error
}

type ErrValidatorHasWrongFormat struct {
    Type, Cause string
}

func (e ErrValidatorHasWrongFormat) Error() string {
    return fmt.Sprintf("Validator for type '%s' has a wrong format: %s", e.Type, e.Cause)
}