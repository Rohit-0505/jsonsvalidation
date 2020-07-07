package jsonschema

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//Dclare the mandatory formatchecker type of struct
type MandatoryFormatChecker struct{}

func validateJSON() {
	gojsonschema.FormatCheckers.Add("mandatory", MandatoryFormatChecker{})

	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/rohitk/Documents/practiceProject/jsonsvalidation/host_schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file:///Users/rohitk/Documents/practiceProject/jsonsvalidation/host_document.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err.Error())
	}
	if result.Valid() {
		fmt.Printf("Validation Success ! The document is valid !!\n")
	} else {
		fmt.Printf("Validation Failed ! The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("ERROR- %s\n", desc)
		}

	}
}

// Define the mandarory format checker
func (f MandatoryFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if ok == false {
		return false
	}
	if len(asString) > 0 {
		return true
	}
	return false
}
