package domain

import (
    "github.com/xeipuuv/gojsonschema"
    "fmt"
)

const VersionNotAssignedYet = -1

type Validator struct {
    Type string
    Version int
    Rules []byte
}

func NewValidator(_type string, rules []byte) (Validator, error) {
    err := checkRulesFormat(rules)
    if err != nil {
        return nil, err
    }

    return &Validator{
        Type: _type,
        Version: VersionNotAssignedYet, // It will be set by the
        Rules: rules,
    }
}

type ValidatorFormatChecker interface {

}

func checkRulesFormat(rules []byte) error {
    schemaLoader := gojsonschema.NewStringLoader(string(rules))
    _, err := gojsonschema.NewSchema(schemaLoader)
    if err != nil {
        return err
    }

    return nil
}

type ErrValidatorHasWrongFormat struct {
    Type, Cause string
}

func (e ErrValidatorHasWrongFormat) Error() string {
    return fmt.Sprintf("Validator for type '%s' has a wrong format: %s", e.Type, e.Cause)
}