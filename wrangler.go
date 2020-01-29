package main

import (
	"errors"
	"sort"
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

	// Prepare parameter/response values and substitute into formula output
	var ProcessError error
	var outputDataset ValidationOutputWrapper
	for _, i := range config.Validations {

		i.OutputFormula = i.Formula
		sortByParameterNameLength(i.Parameters)
		for _, k := range i.Parameters {

			k.OffsetPeriod, ProcessError = GetRelativePeriod(config.Period, k.PeriodOffset, config.Periodicity)
			if ProcessError != nil {
				return ValidationOutputWrapper{}, ProcessError
			}

			k.ReplacementValue, ProcessError = getLookupValue(k, config.Responses, config.Contributors, i.Default)
			if ProcessError != nil {
				return ValidationOutputWrapper{}, ProcessError
			}
			// Substitute our found parameter value into the formula
			i.OutputFormula = strings.ReplaceAll(i.OutputFormula, k.Name, k.ReplacementValue)
			// Replace ' (i.e. single quote) with \" to help simplify json issues
			i.OutputFormula = strings.ReplaceAll(i.OutputFormula, "'", "\"")

		}
		outputDataset.Output = append(outputDataset.Output, ValidationOutput{Formula: i.OutputFormula, Metadata: Metadata{Validation: i.Rule, Reference: config.Reference, Survey: config.Survey, Period: config.Period, ValidationID: i.ValidationID, BpmID: config.BpmID}})
	}

	return outputDataset, nil
}

// getLookupValue ... Work out which value we will use to replace in the validation formula
func getLookupValue(parameters ValidationParameter, responses []Response, contributors []Contributor, defaultValue string) (string, error) {
	switch parameters.Source {
	case "":
		return parameters.Value, nil // Use the provided value with the parameter
	case "response":
		return lookupResponseValue(responses, parameters, defaultValue)
	case "contributor":
		return lookupContributorValue(contributors, parameters, defaultValue)
	}
	return "", errors.New("Unable to lookup value for given source: " + parameters.Source)
}

// lookupResponseValue - Get the response value for a given question
// Assumes the same reference & survey across all data provided
// Does not process repeating data
func lookupResponseValue(responses []Response, vp ValidationParameter, defaultValue string) (string, error) {
	for _, r := range responses {
		if (r.Question == vp.Value) && (r.Period == vp.OffsetPeriod) && (r.Instance == 0) {
			var foundResponse = strings.TrimSpace(r.Response)
			if foundResponse != "" {
				return foundResponse, nil
			}
		}
	}
	return defaultValue, nil

}

// lookupContributorValue - Get the underlying attribute value for the contributor
// Currently hard coded into a switch statement. If required we could make this dynamic
// using reflection, but that will increase complexity
func lookupContributorValue(contributors []Contributor, vp ValidationParameter, defaultValue string) (string, error) {
	for _, ct := range contributors {
		if ct.Period == vp.OffsetPeriod {
			switch vp.Name {
			case "frosic":
				return ct.Frosic, nil
			case "rusic":
				return ct.Frosic, nil
			case "status":
				return ct.Status, nil
			}
		}
	}
	return defaultValue, errors.New("Unable to find contributor value: " + vp.Name + ":" + vp.OffsetPeriod)
}

func sortByParameterNameLength(parameters []ValidationParameter) {
	sort.SliceStable(parameters, func(i, j int) bool {
		return len(parameters[i].Name) >= len(parameters[j].Name)
	})
}
