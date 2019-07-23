package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestWrangle(t *testing.T) {

	start := time.Now()
	var config Config
	err := json.Unmarshal(testJSON, &config)
	if err == nil {
		t.Errorf("Error: %v", err)
	}

	output, err := Wrangle(config)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	elapsed := time.Since(start)
	fmt.Println(output)
	fmt.Printf("Runtime: %s", elapsed)
}

// func TestLookup (t* testing.T){
// 	getLookupValue()
// }

var testJSON = []byte(`
{
"validation_period":"201912",
"validation_reference":"70001",
"validation_survey":"003",
"periodicity":"quarterly",
"contributor":	[
		{"reference":"70001", "period":"201912", "survey":"003", "status":"Clear", "frosic":"92341", "rusic":"92340"},
	],
"response":	[
		{"reference":"70001","period":"201909","survey":"003","instance":0,"question":"3212","response":"16.1"},
		{"reference":"70001","period":"201912","survey":"003","instance":0,"question":"3212","response":"1234.5678"},
		{"reference":"70001","period":"201909","survey":"003","instance":0,"question":"3213","response":"0"},
		{"reference":"70001","period":"201912","survey":"003","instance":0,"question":"3213","response":"12.3"},
		{"reference":"70001","period":"201912","survey":"003","instance":0,"question":"3214","response":"Toast"},
		{"reference":"70001","period":"201912","survey":"003","instance":0,"question":"3215","response":""},
	],
"question_schema": [
		{ "survey":"003", "period":"201909", "question":"3212", "datatype":"numeric", "repeating":false },
		{ "survey":"003", "period":"201912", "question":"3212", "datatype":"numeric", "repeating":false },
		{ "survey":"003", "period":"201909", "question":"3213", "datatype":"numeric", "repeating":false },
		{ "survey":"003", "period":"201912", "question":"3213", "datatype":"numeric", "repeating":false },
		{ "survey":"003", "period":"201909", "question":"3214", "datatype":"text", "repeating":false },
		{ "survey":"003", "period":"201912", "question":"3214", "datatype":"text", "repeating":false },
		{ "survey":"003", "period":"201909", "question":"3215", "datatype":"text", "repeating":false },
		{ "survey":"003", "period":"201912", "question":"3215", "datatype":"text", "repeating":false },
		{ "survey":"003", "period":"201909", "question":"3216", "datatype":"date", "repeating":false },
		{ "survey":"003", "period":"201912", "question":"3216", "datatype":"date", "repeating":false },
	],
"validation_config": [
		{
			"template": "Value present",
			"formula": "question != ''",
			"question_details": [
					{
						"primary_question": "3212",
						"default": "",
						"parameters": [{ "name": "question", "value": "3212", "source": "response", "response_offset": 0 }]
					},
					{
						"primary_question": "3214",
						"default": "",
						"parameters": [{ "name": "question", "value": "3214", "source": "response", "response_offset": 0 }]
					},
					{
						"primary_question": "3215",
						"default": "",
						"parameters": [{ "name": "question", "value": "3215", "source": "response", "response_offset": 0 }]
					}
				]
		},
		{
			"template": "Question v Question",
			"formula": "question != comparison_question",
			"question_details": 
				[
					{
						"primary_question": "3212",
						"default": "0",
						"parameters":
							[
								{ "name": "comparison_question", "value": "3212", "source": "response", "response_offset": 1 },
								{ "name": "question", "value": "3212", "source": "response", "response_offset": 0 }
							]
					},
					{
						"primary_question": "3213",
						"default": "0",
						"parameters": 
							[
								{ "name": "comparison_question", "value": "3213", "source": "response", "response_offset": 1 },
								{ "name": "question", "value": "3213", "source": "response", "response_offset": 0 }
							]
					},
					{
						"primary_question": "3214",
						"default": "0",
						"parameters":
							[
								{ "name": "comparison_question", "value": "3214", "source": "response", "response_offset": 1 },
								{ "name": "question", "value": "3214", "source": "response", "response_offset": 0 }
							]
					}
				]
		},
		{
			"template": "Compulsory Value",
			"formula": "question = 0 OR 'question' = ''",
			"question_details": 
				[
					{
						"primary_question": "3214",
						"default": "0",
						"parameters":
							[
								{ "name": "question", "value": "3214", "source": "response", "response_offset": 0 }
							]
					}
				]
		},
		{
			"template": "Value Present - SIC Exclusion",
			"formula": "question != '' AND 'frosic' NOT IN (range)",
			"question_details": 
				[
					{
						"primary_question": "3212",
						"default": "0",
						"parameters":
							[
								{ "name": "frosic", "value": "", "source": "contributor", "response_offset": 0 },
								{ "name": "question", "value": "3212", "source": "response", "response_offset": 0 },
								{ "name": "range", "value": "'01000','02000','03000'", "source": "", "response_offset": 0 }
							]
					}
				]
		},
		{
			"template": "Period on Period Movement",
			"formula": "(((( question - comparison_question ) / question) > IncreasePercentThreshold AND abs(question - comparison_question) > IncreaseAbsoluteThreshold) OR (((( comparison_question - question ) / comparison_question) > DecreasePercentThreshold AND abs(question-comparison_question) > DecreaseAbsoluteThreshold))",
			"question_details": 
				[
					{
						"primary_question": "3212",
						"default": "0",
						"parameters":
							[
								{ "name": "comparison_question", "value": "3212", "source": "response", "response_offset": 1 },
								{ "name": "question", "value": "3212", "source": "response", "response_offset": 0 },
                                { "name": "IncreasePercentThreshold", "value": "5", "source": "", "response_offset": 0 },
                                { "name": "DecreasePercentThreshold", "value": "10", "source": "", "response_offset": 0 },
								{ "name": "DecreaseAbsoluteThreshold", "value": "100", "source": "", "response_offset": 0 },
                                { "name": "IncreaseAbsoluteThreshold", "value": "200", "source": "", "response_offset": 0 }
							]
					}
				]
		},
		{
			"template": "Period on Period Zero Continuity",
			"formula": "(( question = 0 and comparison_question >= 0 ) OR ( question >= 0 AND comparison_question = 0 )) AND abs(question - comparison_question) > Threshold))",
			"question_details": 
				[
					{
						"primary_question": "3212",
						"default": "0",
						"parameters":
							[
								{ "name": "comparison_question", "value": "3213", "source": "response", "response_offset": 1 },
								{ "name": "question", "value": "3213", "source": "response", "response_offset": 0 },
                                { "name": "Threshold", "value": "25", "source": "", "response_offset": 0 }
							]
					}
				]
		},
		{
			"template": "Question v Threshold",
			"formula": "question operator threshold",
			"question_details": 
				[
					{
						"primary_question": "3212",
						"default": "0",
						"parameters":
							[
								{ "name": "operator", "value": "!=", "source": "", "response_offset": 0 },
								{ "name": "threshold", "value": "335", "source": "", "response_offset": 0 },
								{ "name": "question", "value": "3212", "source": "response", "response_offset": 0 }
							]
					},
					{
						"primary_question": "3213",
						"default": "0",
						"parameters": 
							[
								{ "name": "operator", "value": "<", "source": "", "response_offset": 0 },
								{ "name": "threshold", "value": "26", "source": "", "response_offset": 0 },
								{ "name": "question", "value": "3213", "source": "response", "response_offset": 0 }
							]
					},
					{
						"primary_question": "3214",
						"default": "0",
						"parameters":
							[
								{ "name": "operator", "value": ">=", "source": "", "response_offset": 0 },
								{ "name": "threshold", "value": "965", "source": "", "response_offset": 0 },
								{ "name": "question", "value": "3214", "source": "response", "response_offset": 0 }
							]
					}
				]
		}		
	]		
}
`)
