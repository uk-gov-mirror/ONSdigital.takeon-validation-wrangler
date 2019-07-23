package main

import (
	"errors"
	"strconv"
	"strings"
)

// PeriodFrequency - Holds how often a period changes, i.e. what is the difference between one IDBR period and another.
//   e.g. a monthly survey will change once every month and is every year, a quarterly will change once every 3 months
type PeriodFrequency struct {
	MonthFrequency int
	YearFrequency  int
}

// GetPeriodFrequency - Set how many months/years to change by if we want the previous period
func GetPeriodFrequency(periodicity string) (PeriodFrequency, error) {

	var periodFrequency PeriodFrequency

	switch strings.ToLower(periodicity) {
	case "monthly":
		periodFrequency.MonthFrequency = 1
		periodFrequency.YearFrequency = 1
	case "quarterly":
		periodFrequency.MonthFrequency = 3
		periodFrequency.YearFrequency = 1
	case "biannual":
		periodFrequency.MonthFrequency = 6
		periodFrequency.YearFrequency = 1
	case "annual":
		periodFrequency.MonthFrequency = 12
		periodFrequency.YearFrequency = 1
	case "biennial":
		periodFrequency.MonthFrequency = 12
		periodFrequency.YearFrequency = 2
	case "triennial":
		periodFrequency.MonthFrequency = 12
		periodFrequency.YearFrequency = 3
	default:
		return periodFrequency, errors.New("Unknown periodicity: " + periodicity)
	}

	return periodFrequency, nil
}

// CalculatePreviousPeriod - For a given periodicity/frequency calculate the older IDBR period given an IDBR period as a starting point.
func CalculatePreviousPeriod(pf PeriodFrequency, period string) (string, error) {

	currentYear, err := strconv.Atoi(period[0:4])
	if err != nil {
		return "", errors.New("Unable to extract month from given period: " + period)
	}

	currentMonth, err := strconv.Atoi(period[len(period)-2:])
	if err != nil {
		return "", errors.New("Unable to extract year from given period: " + period)
	}

	previousMonth := currentMonth - pf.MonthFrequency
	previousYear := currentYear

	if previousMonth < 1 {
		previousMonth += 12
		previousYear -= pf.YearFrequency
	}

	previousPeriod := strconv.Itoa(previousYear) + PaddedMonth(previousMonth)
	return previousPeriod, nil

}

// PaddedMonth ... Add a leading zero to any single digit months in line with IDBR period rules
func PaddedMonth(month int) string {
	if month < 10 {
		return "0" + strconv.Itoa(month)
	}
	return strconv.Itoa(month)
}

// GetRelativePeriod ...  assumes offset is >= 0
func GetRelativePeriod(basePeriod string, offset int, periodicity string) (string, error) {

	periodFrequency, err := GetPeriodFrequency(periodicity)
	if err != nil {
		return "", err
	}

	offsetPeriod := basePeriod
	for i := 1; i <= offset; i++ {
		offsetPeriod, err = CalculatePreviousPeriod(periodFrequency, offsetPeriod)
		if err != nil {
			return "", err
		}
	}

	return offsetPeriod, nil
}
