package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWrangle(t *testing.T) {

	start := time.Now()
	var config Config
	err := json.Unmarshal(testJSON, &config)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	output, err := Wrangle(config)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected, _ := Unmarshal(expectedOutputJSON)

	DataToOutput, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("An error occured while marshaling DataToOutput: %s", err)
	}

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	actual, _ := Unmarshal(DataToOutput)
	assert.True(t, Equal(expected, actual))

	elapsed := time.Since(start)
	// fmt.Println(output)
	fmt.Printf("Runtime: %s", elapsed)

}

// func TestLookup (t* testing.T){
// 	getLookupValue()
// }

var testJSON = []byte(`
{
	"validation_period": "201912",
	"validation_reference": "70001",
	"validation_survey": "003",
	"periodicity": "quarterly",
	"bpmid": "",
	"contributor": [
	  {
		"reference": "70001",
		"period": "201912",
		"survey": "003",
		"status": "Clear",
		"frosic": "92341",
		"rusic": "92340"
	  },
	  {
		"reference": "70001",
		"period": "201909",
		"survey": "003",
		"status": "Form Sent Out",
		"frosic": "92341",
		"rusic": "92340"
	  }
	],
	"response": [
	  {
		"reference": "70001",
		"period": "201912",
		"survey": "003",
		"instance": 0,
		"question": "3212",
		"response": "343.3"
	  },
	  {
		"reference": "70001",
		"period": "201912",
		"survey": "003",
		"instance": 0,
		"question": "3213",
		"response": "12.3"
	  },
	  {
		"reference": "70001",
		"period": "201912",
		"survey": "003",
		"instance": 0,
		"question": "3215",
		"response": "Toast"
	  },
	  {
		"reference": "70001",
		"period": "201909",
		"survey": "003",
		"instance": 0,
		"question": "3212",
		"response": "16.1"
	  },
	  {
		"reference": "70001",
		"period": "201909",
		"survey": "003",
		"instance": 0,
		"question": "3213",
		"response": "0"
	  },
	  {
		"reference": "70001",
		"period": "201912",
		"survey": "003",
		"instance": 0,
		"question": "3214",
		"response": "340"
	  },
	  {
		"reference": "70001",
		"period": "201909",
		"survey": "003",
		"instance": 0,
		"question": "3214",
		"response": "125"
	  }
	],
	"question_schema": [
	  {
		"survey": "003",
		"period": "201909",
		"question": "3212",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201909",
		"question": "3213",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201909",
		"question": "3214",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201909",
		"question": "3215",
		"datatype": "text",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201909",
		"question": "3216",
		"datatype": "date",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3212",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3213",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3214",
		"datatype": "numeric",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3215",
		"datatype": "text",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3216",
		"datatype": "date",
		"repeating": false
	  },
	  {
		"survey": "003",
		"period": "201912",
		"question": "3217",
		"datatype": "date",
		"repeating": false
	  }
	],
	"validation_config": [
	  {
		"template": "Value present",
		"formula": "question != ''",
		"question_details": [
		  { 
			"validationid": 23412,
			"primary_question": "3212",
			"default": "",
			"parameters": [
			  {
				"name": "question",
				"value": "3212",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  { "validationid": 23413,
			"primary_question": "3213",
			"default": "",
			"parameters": [
			  {
				"name": "question",
				"value": "3213",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  { 
			"validationid": 23414,
			"primary_question": "3214",
			"default": "",
			"parameters": [
			  {
				"name": "question",
				"value": "3214",
				"source": "response",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Question v Question",
		"formula": "question != comparison_question",
		"question_details": [
		  {
			"validationid": 23412,
			"primary_question": "3212",
			"default": "0",
			"parameters": [
			  {
				"name": "comparison_question",
				"value": "3212",
				"source": "response",
				"period_offset": 1
			  },
			  {
				"name": "question",
				"value": "3212",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  {
			"validationid": 23413,
			"primary_question": "3213",
			"default": "0",
			"parameters": [
			  {
				"name": "comparison_question",
				"value": "3213",
				"source": "response",
				"period_offset": 1
			  },
			  {
				"name": "question",
				"value": "3213",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  {
			"validationid": 23414,
			"primary_question": "3214",
			"default": "0",
			"parameters": [
			  {
				"name": "comparison_question",
				"value": "3214",
				"source": "response",
				"period_offset": 1
			  },
			  {
				"name": "question",
				"value": "3214",
				"source": "response",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Compulsory Value",
		"formula": "question = 0 OR 'question' = ''",
		"question_details": [
		  { 
			"validationid": 23414,
			"primary_question": "3214",
			"default": "0",
			"parameters": [
			  {
				"name": "question",
				"value": "3214",
				"source": "response",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Value Present - SIC Exclusion",
		"formula": "question != '' AND 'frosic' NOT IN (range)",
		"question_details": [
		  {
			"validationid": 23412,
			"primary_question": "3212",
			"default": "0",
			"parameters": [
			  {
				"name": "frosic",
				"value": "",
				"source": "contributor",
				"period_offset": 0
			  },
			  {
				"name": "question",
				"value": "3212",
				"source": "response",
				"period_offset": 0
			  },
			  {
				"name": "range",
				"value": "'01000','02000','03000'",
				"source": "",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Period on Period Movement",
		"formula": "(((( question - comparison_question ) / question) > IncreasePercentThreshold AND abs(question - comparison_question) > IncreaseAbsoluteThreshold) OR (((( comparison_question - question ) / comparison_question) > DecreasePercentThreshold AND abs(question-comparison_question) > DecreaseAbsoluteThreshold))",
		"question_details": [
		  {
			"validationid": 23412,
			"primary_question": "3212",
			"default": "0",
			"parameters": [
			  {
				"name": "comparison_question",
				"value": "3212",
				"source": "response",
				"period_offset": 1
			  },
			  {
				"name": "question",
				"value": "3212",
				"source": "response",
				"period_offset": 0
			  },
			  {
				"name": "IncreasePercentThreshold",
				"value": "5",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "DecreasePercentThreshold",
				"value": "10",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "DecreaseAbsoluteThreshold",
				"value": "100",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "IncreaseAbsoluteThreshold",
				"value": "200",
				"source": "",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Period on Period Zero Continuity",
		"formula": "(( question = 0 and comparison_question >= 0 ) OR ( question >= 0 AND comparison_question = 0 )) AND abs(question - comparison_question) > Threshold))",
		"question_details": [
		  { 
			"validationid": 23412,
			"primary_question": "3212",
			"default": "0",
			"parameters": [
			  {
				"name": "comparison_question",
				"value": "3213",
				"source": "response",
				"period_offset": 1
			  },
			  {
				"name": "question",
				"value": "3213",
				"source": "response",
				"period_offset": 0
			  },
			  {
				"name": "Threshold",
				"value": "25",
				"source": "",
				"period_offset": 0
			  }
			]
		  }
		]
	  },
	  {
		"template": "Question v Threshold",
		"formula": "question operator threshold",
		"question_details": [
		  { 
			"validationid": 23412,
			"primary_question": "3212",
			"default": "0",
			"parameters": [
			  {
				"name": "operator",
				"value": "!=",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "threshold",
				"value": "335",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "question",
				"value": "3212",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  {
			"validationid": 23413,
			"primary_question": "3213",
			"default": "0",
			"parameters": [
			  {
				"name": "operator",
				"value": "<",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "threshold",
				"value": "26",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "question",
				"value": "3213",
				"source": "response",
				"period_offset": 0
			  }
			]
		  },
		  {
			"validationid": 23414,
			"primary_question": "3214",
			"default": "0",
			"parameters": [
			  {
				"name": "operator",
				"value": ">=",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "threshold",
				"value": "965",
				"source": "",
				"period_offset": 0
			  },
			  {
				"name": "question",
				"value": "3214",
				"source": "response",
				"period_offset": 0
			  }
			]
		  }
		]
	  }
	]
  }
`)

var expectedOutputJSON = []byte(`
{
    "validation_input": [
        {
            "formula": "343.3 != \"\"",
            "metadata": {
                "validation": "Value present",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "12.3 != \"\"",
            "metadata": {
                "validation": "Value present",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23413,
                "bpmid": ""
            }
        },
        {
            "formula": "340 != \"\"",
            "metadata": {
                "validation": "Value present",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23414,
                "bpmid": ""
            }
        },
        {
            "formula": "343.3 != 16.1",
            "metadata": {
                "validation": "Question v Question",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "12.3 != 0",
            "metadata": {
                "validation": "Question v Question",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23413,
                "bpmid": ""
            }
        },
        {
            "formula": "340 != 125",
            "metadata": {
                "validation": "Question v Question",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23414,
                "bpmid": ""
            }
        },
        {
            "formula": "340 = 0 OR \"340\" = \"\"",
            "metadata": {
                "validation": "Compulsory Value",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23414,
                "bpmid": ""
            }
        },
        {
            "formula": "343.3 != \"\" AND \"92341\" NOT IN (\"01000\",\"02000\",\"03000\")",
            "metadata": {
                "validation": "Value Present - SIC Exclusion",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "(((( 343.3 - 16.1 ) / 343.3) > 5 AND abs(343.3 - 16.1) > 200) OR (((( 16.1 - 343.3 ) / 16.1) > 10 AND abs(343.3-16.1) > 100))",
            "metadata": {
                "validation": "Period on Period Movement",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "(( 12.3 = 0 and 0 >= 0 ) OR ( 12.3 >= 0 AND 0 = 0 )) AND abs(12.3 - 0) > 25))",
            "metadata": {
                "validation": "Period on Period Zero Continuity",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "343.3 != 335",
            "metadata": {
                "validation": "Question v Threshold",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23412,
                "bpmid": ""
            }
        },
        {
            "formula": "12.3 < 26",
            "metadata": {
                "validation": "Question v Threshold",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23413,
                "bpmid": ""
            }
        },
        {
            "formula": "340 >= 965",
            "metadata": {
                "validation": "Question v Threshold",
                "reference": "70001",
                "period": "201912",
                "survey": "003",
                "validationid": 23414,
                "bpmid": ""
            }
        }
    ]
}
`)
