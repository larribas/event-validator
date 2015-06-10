package formats

import (
	"github.com/xeipuuv/gojsonschema"
)

type JSONSchemaFormatChecker struct{}

func (c *JSONSchemaFormatChecker) Check(rules []byte) error {
	schemaLoader := gojsonschema.NewStringLoader(string(rules))
	_, err := gojsonschema.NewSchema(schemaLoader)
	return err
}
