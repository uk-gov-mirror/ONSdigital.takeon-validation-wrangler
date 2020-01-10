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
	"period":"201801",
	"reference":"12345678001",
	"survey":"999A",
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
	"validation_config": [
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "\"question\" != \"\"",
            "validationid": 10,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "VP",
            "defaultvalue": "",
            "primaryquestion": "3000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Value Present",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "VP",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlZQIiwwXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDEwXQ==",
            "parameters": [
                {
                    "validationid": 10,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMTAsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "3000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "E",
            "baseformula": "\"question\" != \"\"",
            "validationid": 11,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "VP",
            "defaultvalue": "",
            "primaryquestion": "2000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Value Present",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "VP",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlZQIiwwXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDExXQ==",
            "parameters": [
                {
                    "validationid": 11,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMTEsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "2000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "\"question\" != \"\"",
            "validationid": 12,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "VP",
            "defaultvalue": "",
            "primaryquestion": "1000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Value Present",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "VP",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlZQIiwwXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDEyXQ==",
            "parameters": [
                {
                    "validationid": 12,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMTIsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "\"question\" != \"\"",
            "validationid": 13,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "VP",
            "defaultvalue": "",
            "primaryquestion": "4000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Value Present",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "VP",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlZQIiwwXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDEzXQ==",
            "parameters": [
                {
                    "validationid": 13,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMTMsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "4000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "abs(question - comparison_question) > threshold AND question > 0 AND comparison_question > 0",
            "validationid": 20,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "POPM",
            "defaultvalue": "0",
            "primaryquestion": "1000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Period on Period Movement",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPM",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUE0iLDBd"
                },
                {
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPM",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUE0iLDFd"
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDIwXQ==",
            "parameters": [
                {
                    "validationid": 20,
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjAsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 20,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjAsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 20,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "threshold",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjAsIkRlZmF1bHQiLCJEZWZhdWx0IiwidGhyZXNob2xkIl0=",
                    "attributename": "Default",
                    "source": "",
                    "value": "20000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "E",
            "baseformula": "abs(question - comparison_question) > threshold AND question > 0 AND comparison_question > 0",
            "validationid": 21,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "POPM",
            "defaultvalue": "0",
            "primaryquestion": "1001",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Period on Period Movement",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPM",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUE0iLDBd"
                },
                {
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPM",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUE0iLDFd"
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDIxXQ==",
            "parameters": [
                {
                    "validationid": 21,
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjEsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1001"
                },
                {
                    "validationid": 21,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjEsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1001"
                },
                {
                    "validationid": 21,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "threshold",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMjEsIkRlZmF1bHQiLCJEZWZhdWx0IiwidGhyZXNob2xkIl0=",
                    "attributename": "Default",
                    "source": "",
                    "value": "0"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "question != comparison_question",
            "validationid": 30,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "QVDQ",
            "defaultvalue": "0",
            "primaryquestion": "1000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Question vs Derived Question",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "QVDQ",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlFWRFEiLDBd"
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDMwXQ==",
            "parameters": [
                {
                    "validationid": 30,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMzAsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "4000"
                },
                {
                    "validationid": 30,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMzAsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "E",
            "baseformula": "question != comparison_question",
            "validationid": 31,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "QVDQ",
            "defaultvalue": "0",
            "primaryquestion": "1001",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Question vs Derived Question",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "QVDQ",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlFWRFEiLDBd"
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDMxXQ==",
            "parameters": [
                {
                    "validationid": 31,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMzEsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "4001"
                },
                {
                    "validationid": 31,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsMzEsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1001"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "W",
            "baseformula": "question != comparison_question AND ( question = 0 OR comparison_question = 0 ) AND abs(question - comparison_question) > threshold",
            "validationid": 40,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "POPZC",
            "defaultvalue": "0",
            "primaryquestion": "1000",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Period on Period Zero Continuity",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPZC",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUFpDIiwwXQ=="
                },
                {
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPZC",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUFpDIiwxXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDQwXQ==",
            "parameters": [
                {
                    "validationid": 40,
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDAsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 40,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDAsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 40,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "threshold",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDAsIkRlZmF1bHQiLCJEZWZhdWx0IiwidGhyZXNob2xkIl0=",
                    "attributename": "Default",
                    "source": "",
                    "value": "30000"
                }
            ]
        },
        {
            "formid": 1,
            "severity": "E",
            "baseformula": "question != comparison_question AND ( question = 0 OR comparison_question = 0 ) AND abs(question - comparison_question) > threshold",
            "validationid": 41,
            "createddate": "2020-01-10T14:12:40.50121+00:00",
            "rule": "POPZC",
            "defaultvalue": "0",
            "primaryquestion": "1001",
            "createdby": "u0a5c346821b42e",
            "lastupdateddate": null,
            "lastupdatedby": null,
            "name": "Period on Period Zero Continuity",
            "period_offset": [
                {
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPZC",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUFpDIiwwXQ=="
                },
                {
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.498317+00:00",
                    "lastupdateddate": null,
                    "lastupdatedby": null,
                    "rule": "POPZC",
                    "id": "WyJ2YWxpZGF0aW9ucGVyaW9kcyIsIlBPUFpDIiwxXQ=="
                }
            ],
            "id": "WyJ2YWxpZGF0aW9uZm9ybXMiLDQxXQ==",
            "parameters": [
                {
                    "validationid": 41,
                    "periodoffset": 1,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "comparison_question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDEsIkRlZmF1bHQiLCJEZWZhdWx0IiwiY29tcGFyaXNvbl9xdWVzdGlvbiJd",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 41,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "question",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDEsIkRlZmF1bHQiLCJEZWZhdWx0IiwicXVlc3Rpb24iXQ==",
                    "attributename": "Default",
                    "source": "response",
                    "value": "1000"
                },
                {
                    "validationid": 41,
                    "periodoffset": 0,
                    "createdby": "u0a5c346821b42e",
                    "createddate": "2020-01-10T14:12:40.5034+00:00",
                    "lastupdateddate": null,
                    "attributevalue": "Default",
                    "parameter": "threshold",
                    "lastupdatedby": null,
                    "id": "WyJ2YWxpZGF0aW9ucGFyYW1ldGVycyIsNDEsIkRlZmF1bHQiLCJEZWZhdWx0IiwidGhyZXNob2xkIl0=",
                    "attributename": "Default",
                    "source": "",
                    "value": "0"
                }
            ]
        }
	],
	"question_schema": [
        {
            "period": "201801",
            "questioncode": "1000",
            "survey": "999A",
            "type": "NUMERIC"
        },
        {
            "period": "201801",
            "questioncode": "1001",
            "survey": "999A",
            "type": "NUMERIC"
        },
        {
            "period": "201801",
            "questioncode": "2000",
            "survey": "999A",
            "type": "TICKBOX-Yes"
        },
        {
            "period": "201801",
            "questioncode": "3000",
            "survey": "999A",
            "type": "Text"
        },
        {
            "period": "201801",
            "questioncode": "4000",
            "survey": "999A",
            "type": "NUMERIC"
        },
        {
            "period": "201801",
            "questioncode": "4001",
            "survey": "999A",
            "type": "NUMERIC"
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
