package main

import (
	"errors"
	"strings"
)

// ValidationOutputWrapper ...
type ValidationOutputWrapper struct {
	Output []ValidationOutput `json:"validation_input"`
}

// ValidationOutput ...
type ValidationOutput struct {
	Formula  string   `json:"formula"`
	Metadata Metadata `json:"metadata"`
}

// Metadata ...
type Metadata struct {
	Validation   string `json:"validation"`
	Reference    string `json:"reference"`
	Period       string `json:"period"`
	Survey       string `json:"survey"`
	ValidationID int    `json:"validationid"`
	BpmID        string `json:"bpmid"`
}

// Wrangle -- main entry point
func Wrangle(config Config) (ValidationOutputWrapper, error) {

	// Ensure response dataset is ‘complete’ using the form definitions
	// completeResponse := GenerateCompleteDataset(config.QuestionSchema, config.Responses)

	// If repeating data is defined:
	// -- If totals are required - generate a totals dataset
	// -- If grouped totals are required – generate a grouped totals dataset

	// Prepare parameter/response values and substitute into formula output
	var ProcessError error
	var outputDataset ValidationOutputWrapper
	for _, i := range config.Validations {
		for _, j := range i.QuestionDetails {
			j.OutputFormula = i.Formula
			for _, k := range j.Parameters {

				k.OffsetPeriod, ProcessError = GetRelativePeriod(config.Period, k.ResponseOffset, config.Periodicity)
				if ProcessError != nil {
					return ValidationOutputWrapper{}, ProcessError
				}

				k.ReplacementValue, ProcessError = getLookupValue(k, config.Responses, config.Contributors)
				if ProcessError != nil {
					return ValidationOutputWrapper{}, ProcessError
				}
				// Substitute our found parameter value into the formula
				j.OutputFormula = strings.ReplaceAll(j.OutputFormula, k.Name, k.ReplacementValue)

			}
			outputDataset.Output = append(outputDataset.Output, ValidationOutput{Formula: j.OutputFormula, Metadata: Metadata{Validation: i.Template, Reference: config.Reference, Survey: config.Survey, Period: config.Period, ValidationID: j.ValidationID, BpmID: config.BpmID}})
		}
	}

	return outputDataset, nil
}

// getLookupValue ... Work out which value we will use to replace in the validation formula
func getLookupValue(parameters ValidationParameter, responses []Response, contributors []Contributor) (string, error) {
	switch parameters.Source {
	case "":
		return parameters.Value, nil // Use the provided value with the parameter
	case "response":
		return lookupResponseValue(responses, parameters)
	case "contributor":
		return lookupContributorValue(contributors, parameters)
	}
	return "", errors.New("Unable to lookup value for given source: " + parameters.Source)
}

// lookupResponseValue - Get the response value for a given question
// Assumes the same reference period & survey across all data provided
// Does not process repeating data
func lookupResponseValue(responses []Response, vp ValidationParameter) (string, error) {
	for _, r := range responses {
		if (r.Question == vp.Value) && (r.Period == vp.OffsetPeriod) && (r.Instance == 0) {
			return r.Response, nil // Assumes all responses have same reference & survey
		}
	}
	return "", errors.New("Unable to find response value for given response: " + vp.Name + ":" + vp.OffsetPeriod + ":" + vp.Value)
}

// lookupContributorValue - Get the underlying attribute value for the contributor
// Currently hard coded into a switch statement. If required we could make this dynamic
// using reflection, but that will increase complexity
func lookupContributorValue(contributors []Contributor, vp ValidationParameter) (string, error) {
	for _, ct := range contributors {
		if ct.Period == vp.OffsetPeriod {
			switch vp.Name {
			case "frosic":
				return ct.Frosic, nil
			case "rusic":
				return ct.Frosic, nil
			}
		}
	}
	return "", errors.New("Unable to find contributor value: " + vp.Name + ":" + vp.OffsetPeriod)
}
