package formats

import (
    "testing"
)

func TestRightValidatorFormat(t *testing.T) {
    checker := &JSONSchemaFormatChecker{}
    err := checker.Check([]byte(`{"type": "string"}`))
    if err != nil {
        t.Errorf("Expected JSONSchemaFormatChecker::Check not to return an error for a validator with the right format. Instead it returned '%s'", err)
    }
}

func TestWrongValidatorFormat(t *testing.T) {
    checker := &JSONSchemaFormatChecker{}
    err := checker.Check([]byte(`some random string that does not represent a json schema`))
    if err == nil {
        t.Error("Expected JSONSchemaFormatChecker::Check to return an error for a validator with a wrong format")
    }
}