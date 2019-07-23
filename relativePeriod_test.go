package main

import (
	"strings"
	"testing"
)

// ErrorContains checks if the error message in out contains the text in want.
// This is safe when out is nil. Use an empty string for want if you want to test that err is nil.
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}

func TestGetPeriodFrequency(t *testing.T) {

	tests := []struct {
		periodicity      string
		monthlyFrequency int
		yearlyFrequency  int
		expectedError    string
	}{
		{"monthly", 1, 1, ""},
		{"annual", 12, 1, ""},
		{"quarterly", 3, 1, ""},
		{"biennial", 12, 2, ""},
		{"triennial", 12, 3, ""},
		{"biannual", 6, 1, ""},
		{"Beehive", 0, 0, "Unknown periodicity: Beehive"},
	}

	for _, test := range tests {
		frequency, err := GetPeriodFrequency(test.periodicity)

		if (frequency.MonthFrequency != test.monthlyFrequency) || (frequency.YearFrequency != test.yearlyFrequency) || (!ErrorContains(err, test.expectedError)) {
			t.Errorf("Incorrect frequencies, expected Monthly:%d Yearly:%d, Error: %v, got: Monthly: %d, Yearly: %d, Error: %v", test.monthlyFrequency, test.yearlyFrequency, test.expectedError, frequency.MonthFrequency, frequency.YearFrequency, err)
		}
	}
}

func TestPaddedMonth(t *testing.T) {

	tests := []struct {
		month       int
		paddedMonth string
	}{
		{1, "01"},
		{2, "02"},
		{3, "03"},
		{4, "04"},
		{5, "05"},
		{6, "06"},
		{7, "07"},
		{8, "08"},
		{9, "09"},
		{10, "10"},
		{11, "11"},
		{12, "12"},
	}

	for _, test := range tests {
		paddedMonth := PaddedMonth(test.month)
		if paddedMonth != test.paddedMonth {
			t.Errorf("Incorrect padded month --> Expected %v   Got: %v", test.paddedMonth, paddedMonth)
		}
	}
}

func TestCalculatePreviousPeriod(t *testing.T) {

	tests := []struct {
		currentPeriod  string
		previousPeriod string
		frequency      PeriodFrequency
		expectedError  string
	}{
		{"201912", "201911", PeriodFrequency{1, 1}, ""},
		{"201901", "201812", PeriodFrequency{1, 1}, ""},
		{"201912", "201909", PeriodFrequency{3, 1}, ""},
		{"201003", "200912", PeriodFrequency{3, 1}, ""},
		{"201912", "201906", PeriodFrequency{6, 1}, ""},
		{"202003", "201909", PeriodFrequency{6, 1}, ""},
		{"201912", "201903", PeriodFrequency{9, 1}, ""},
		{"201912", "201812", PeriodFrequency{12, 1}, ""},
		// Check for nil frequencies
		// Check for bad periods
	}

	for _, test := range tests {

		period, err := CalculatePreviousPeriod(test.frequency, test.currentPeriod)

		if (period != test.previousPeriod) || (!ErrorContains(err, test.expectedError)) {
			t.Errorf("Incorrect period --> Expected %v Error: %v got: %v, Error: %v", test.previousPeriod, test.expectedError, period, err)
		}
	}
}

func TestGetRelativePeriod(t *testing.T) {

	tests := []struct {
		basePeriod     string
		offset         int
		periodicity    string
		expectedPeriod string
		expectedError  string
	}{
		{"201012", 0, "annual", "201012", ""},
		{"201012", 1, "annual", "200912", ""},
		{"201012", 2, "annual", "200812", ""},
		{"201012", 10, "annual", "200012", ""},
		{"201112", 0, "quarterly", "201112", ""},
		{"201112", 1, "quarterly", "201109", ""},
		{"201112", 2, "quarterly", "201106", ""},
		{"201112", 3, "quarterly", "201103", ""},
		{"201112", 4, "quarterly", "201012", ""},
		{"201312", 0, "monthly", "201312", ""},
		{"201305", 1, "monthly", "201304", ""},
		{"201303", 3, "monthly", "201212", ""},
		{"201311", 12, "monthly", "201211", ""},
	}

	for _, test := range tests {

		period, err := GetRelativePeriod(test.basePeriod, test.offset, test.periodicity)

		if (period != test.expectedPeriod) || (!ErrorContains(err, test.expectedError)) {
			t.Errorf("Incorrect period --> Expected %v Error: %v got: %v, Error: %v", test.expectedPeriod, test.expectedError, period, err)
		}
	}
}
