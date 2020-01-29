package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrangling(t *testing.T) {

	tests := []struct {
		jsonConfig     []byte
		expectedConfig []byte
	}{
		{testParametersPopulatedAcrossPeriodsJSON, testOutputParametersPopulatedAcrossPeriodsJSON},
		{testWhitespaceInResponseJSON, expectedWhitespaceInResponseJSON},
	}

	var expected ValidationOutputWrapper
	var config Config

	for _, test := range tests {
		// Prepare data
		json.Unmarshal(test.jsonConfig, &config)
		json.Unmarshal(test.expectedConfig, &expected)
		var actual, _ = Wrangle(config)
		assert.EqualValues(t, actual, expected)
	}
}

func BenchmarkWrangling(b *testing.B) {

	for i := 0; i < b.N; i++ {
		var config Config
		json.Unmarshal(testParametersPopulatedAcrossPeriodsJSON, &config)
		output, _ := Wrangle(config)
		json.Marshal(output)
	}

}

var testParametersPopulatedAcrossPeriodsJSON = []byte(`
{
	"period":"202101",
	"reference":"12345678001",
	"survey":"999A",
	"periodicity":"Monthly",
	"bpmid":"Dummy",
	"contributor":[
	   	{ "reference":"12345678001", "period":"202101", "survey":"999A", "status":"Validations Triggered", "frozensic":"90901", "rusic":"90902", "turnover":"4654"},
		{ "reference":"12345678001", "period":"202012", "survey":"999A", "status":"Validations Triggered", "frozensic":"90901", "rusic":"90902", "turnover":"4654"}
	],
	"response":[
	   	{ "reference":"12345678001", "period":"202101", "survey":"999A", "questioncode":"1000", "instance":0, "response":"54321" },
		{ "reference":"12345678001", "period":"202012", "survey":"999A", "questioncode":"1001", "instance":0, "response":"12345" }
	],
	"validation_config": [
        {
			"rule": "RAR",
			"baseformula": "abs(question - comparison_question - missing_question) > threshold + frosic AND 'status' = 'Clear'",
			"validationid": 1,
            "defaultvalue": "150.53",
            "primaryquestion": "5000",
            "parameters": [
				{ "parameter": "question", "value": "1000", "source": "response", "periodoffset": 0 },
				{ "parameter": "comparison_question", "value": "1001", "source": "response", "periodoffset": 1 },
				{ "parameter": "missing_question", "value": "1004", "source": "response", "periodoffset": 4 },
				{ "parameter": "threshold", "value": "1000", "source": "", "periodoffset": 0 },
				{ "parameter": "frosic", "value": "", "source": "contributor", "periodoffset": 0 },
				{ "parameter": "status", "value": "", "source": "contributor", "periodoffset": 0 }
            ]
        }
    ]
 }
`)

var testOutputParametersPopulatedAcrossPeriodsJSON = []byte(`
{
    "validation_input": [{
		"formula": "abs(54321 - 12345 - 150.53) > 1000 + 90901 AND \"Validations Triggered\" = \"Clear\"",
		"metadata": {"validation": "RAR", "reference": "12345678001", "period": "202101", "survey": "999A", "validationid": 1, "bpmid": "Dummy"}
        }
    ]
}
`)

var testWhitespaceInResponseJSON = []byte(`
{
	"period":"202101",
	"reference":"12345678001",
	"survey":"999A",
	"periodicity":"Monthly",
	"bpmid":"Dummy",
	"contributor":[
	   	{ "reference":"12345678001", "period":"202101", "survey":"999A", "status":"Validations Triggered", "frozensic":"90901", "rusic":"90902", "turnover":"4654"}
	],
	"response":[
	   	{ "reference":"12345678001", "period":"202101", "survey":"999A", "questioncode":"1000", "instance":0, "response":" " }
	],
	"validation_config": [
        {
			"rule": "testWhiteSpace",
			"baseformula": "'question' != ''",
			"validationid": 2,
            "defaultvalue": "",
            "primaryquestion": "1000",
            "parameters": [
				{ "parameter": "question", "value": "1000", "source": "response", "periodoffset": 0 }
            ]
        }
    ]
 }
`)

var expectedWhitespaceInResponseJSON = []byte(`
{
    "validation_input": [{
		"formula": "\"\" != \"\"",
		"metadata": {"validation": "testWhiteSpace", "reference": "12345678001", "period": "202101", "survey": "999A", "validationid": 2, "bpmid": "Dummy"}
        }
    ]
}
`)
