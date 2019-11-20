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
	"validation_period":"201801",
	"validation_reference":"12345678001",
	"validation_survey":"999A",
	"periodicity":"Monthly",
	"bpmid":"Dummy",
	"contributor":[ 
	   { 
		  "formid":1,
		  "birthdate":"",
		  "selectiontype":" ",
		  "createddate":"2019-10-25T10:37:51.602379+00:00",
		  "checkletter":" ",
		  "rusicoutdated":"     ",
		  "tradingstyle":"",
		  "frozenemployees":"0",
		  "companyregistrationnumber":"",
		  "reference":"12345678001",
		  "reportingunitmarker":" ",
		  "inclusionexclusion":" ",
		  "legalstatus":" ",
		  "createdby":"fisdba",
		  "lastupdateddate":null,
		  "rusic":"     ",
		  "contact":"",
		  "lastupdatedby":null,
		  "frozenturnover":"0",
		  "currency":"S",
		  "receiptdate":"2019-10-25T10:37:51.602379+00:00",
		  "frozensicoutdated":"     ",
		  "fax":"",
		  "frozenemployment":"0",
		  "turnover":"0",
		  "payereference":"",
		  "period":"201801",
		  "wowenterprisereference":"",
		  "numberlivevat":"0",
		  "telephone":"",
		  "employment":"0",
		  "numberlivepaye":"0",
		  "vatreference":"",
		  "lockedby":null,
		  "frozenfteemployment":"0.000",
		  "cellnumber":0,
		  "fteemployment":"0.000",
		  "lockeddate":null,
		  "survey":"999A",
		  "enterprisereference":"          ",
		  "numberlivelocalunits":"0",
		  "employees":"0",
		  "region":"  ",
		  "frozensic":"     ",
		  "status":"Validations Triggered"
	   }
	],
	"response":[ 
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"45000",
		  "questioncode":"1000",
		  "survey":"999A"
	   },
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"35000",
		  "questioncode":"1001",
		  "survey":"999A"
	   },
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"1",
		  "questioncode":"2000",
		  "survey":"999A"
	   },
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"3",
		  "questioncode":"3000",
		  "survey":"999A"
	   },
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"80000",
		  "questioncode":"4000",
		  "survey":"999A"
	   },
	   { 
		  "reference":"12345678001",
		  "period":"201801",
		  "instance":0,
		  "response":"10000",
		  "questioncode":"4001",
		  "survey":"999A"
	   }
	],
	"validation_config":[ 
	   { 
		  "template":"POPM",
		  "formula":"abs(question - comparison_question) > threshold AND question > 0 AND comparison_question > 0",
		  "question_details":[ 
			 { 
				"validationid":21,
				"primary_question":"1001",
				"default":"0",
				"parameters":[ 
				   { 
					  "periodoffset":1,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"threshold",
					  "source":"",
					  "value":"20000"
				   },
				   { 
					  "periodoffset":1,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"1001"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1001"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"threshold",
					  "source":"",
					  "value":"0"
				   }
				]
			 }
		  ]
	   },
	   { 
		  "template":"POPZC",
		  "formula":"question != comparison_question AND ( question = 0 OR comparison_question = 0 ) AND abs(question - comparison_question) > threshold",
		  "question_details":[ 
			 { 
				"validationid":41,
				"primary_question":"1001",
				"default":"0",
				"parameters":[ 
				   { 
					  "periodoffset":1,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"threshold",
					  "source":"",
					  "value":"30000"
				   },
				   { 
					  "periodoffset":1,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"threshold",
					  "source":"",
					  "value":"0"
				   }
				]
			 }
		  ]
	   },
	   { 
		  "template":"QVDQ",
		  "formula":"question != comparison_question",
		  "question_details":[ 
			 { 
				"validationid":31,
				"primary_question":"1001",
				"default":"0",
				"parameters":[ 
				   { 
					  "periodoffset":0,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"4000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"comparison_question",
					  "source":"response",
					  "value":"4001"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1001"
				   }
				]
			 }
		  ]
	   },
	   { 
		  "template":"VP",
		  "formula":"\"question\" != \"\"",
		  "question_details":[ 
			 { 
				"validationid":13,
				"primary_question":"4000",
				"default":"",
				"parameters":[ 
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"3000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"2000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"1000"
				   },
				   { 
					  "periodoffset":0,
					  "parameter":"question",
					  "source":"response",
					  "value":"4000"
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
            "formula": "abs(45000 - 0) > 20000 AND 45000 > 0 AND 0 > 0",
            "metadata": {
                "validation": "POPM",
                "reference": "12345678001",
                "period": "201801",
                "survey": "999A",
                "validationid": 21,
                "bpmid": "Dummy"
            }
        },
        {
            "formula": "45000 != 0 AND ( 45000 = 0 OR 0 = 0 ) AND abs(45000 - 0) > 30000",
            "metadata": {
                "validation": "POPZC",
                "reference": "12345678001",
                "period": "201801",
                "survey": "999A",
                "validationid": 41,
                "bpmid": "Dummy"
            }
        },
        {
            "formula": "45000 != 80000",
            "metadata": {
                "validation": "QVDQ",
                "reference": "12345678001",
                "period": "201801",
                "survey": "999A",
                "validationid": 31,
                "bpmid": "Dummy"
            }
        },
        {
            "formula": "\"3\" != \"\"",
            "metadata": {
                "validation": "VP",
                "reference": "12345678001",
                "period": "201801",
                "survey": "999A",
                "validationid": 13,
                "bpmid": "Dummy"
            }
        }
    ]
}
`)
