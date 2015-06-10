package formats

import (
	"github.com/xeipuuv/gojsonschema"
)

// JSONSchemaFormatChecker ensures that the given set of rules form a valid JSON Schema
type JSONSchemaFormatChecker struct{}

func (c *JSONSchemaFormatChecker) Check(rules []byte) error {
	schemaLoader := gojsonschema.NewStringLoader(string(rules))
	_, err := gojsonschema.NewSchema(schemaLoader)
	return err
}
