package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	_divisorMap = map[PaidAt]float64{
		PaidAt_Undefined:  0.0,
		PaidAt_Monthly:    12.0,
		PaidAt_Quarterly:  4.0,
		PaidAt_Annually:   1.0,
		PaidAt_AtMaturity: 1.0,
	}
)

// Calculate calculates the final balance of a term deposit after interest.
func Calculate(startDeposit int, interestRate float64, termLength *TermLength, paidAt PaidAt) (float64, error) {
	fmt.Println("Running calculation for final balance...")
	finalBalance, err := calculate(startDeposit, interestRate, termLength, paidAt)

	if err != nil {
		fmt.Println("Final balance calculation unsuccessful.")

		return 0, err
	}

	fmt.Println("Final balance calculation successful.")

	return finalBalance, nil
}

// Validate validates the user inputs to determine whether can be used for
// calculating the final balance of a term deposit.
func Validate(startDeposit int, interestRate float64, termLength, interestPaid string) (*NormalisedInputValues, []error) {
	fmt.Println("Running input validation...")

	normalised, err := validate(startDeposit, interestRate, termLength, interestPaid)
	if err != nil {
		fmt.Println("Input validation unsuccessful.")

		return nil, err
	}

	fmt.Println("Input validation successful.")

	return normalised, nil
}

// validate validates the raw user inputs to see whether they are valid for
// the final balance calculation. If all inputs are valid, a struct containing
// the normalised values in returned, else an error is returned.
func validate(inputStartDeposit int, inputInterestRate float64, inputTermLength, inputInterestPaid string) (*NormalisedInputValues, []error) {
	errors := make([]error, 0)

	// Validate start deposit
	err := validateStartDeposit(inputStartDeposit)
	if err != nil {
		errors = append(errors, &ValidationError{
			err:   err,
			field: "start-deposit",
		})
	}

	// Validate interest rate
	err = validateInterestRate(inputInterestRate)
	if err != nil {
		errors = append(errors, &ValidationError{
			err:   err,
			field: "interest-rate",
		})
	}

	// Validate term length
	termLength, err := validateTermLength(inputTermLength)
	if err != nil {
		errors = append(errors, &ValidationError{
			err:   err,
			field: "term-length",
		})
	}

	// Validate interest paid
	paidAt, err := validateInterestPaid(inputInterestPaid)
	if err != nil {
		errors = append(errors, &ValidationError{
			err:   err,
			field: "interest-paid",
		})
	}

	// If more than one error, validation did not pass
	if len(errors) > 0 {
		return nil, errors
	}

	// Otherwise return a normalised struct of values for calculation
	return &NormalisedInputValues{
		StartDeposit:   inputStartDeposit,
		InterestRate:   inputInterestRate,
		TermLength:     termLength,
		PaidAtInterval: paidAt,
	}, nil
}

// validateStartDeposit Returns true if startDeposit is greater than zero, else returns false and an error
func validateStartDeposit(startDeposit int) error {
	if startDeposit <= 0 {
		return fmt.Errorf("value cannot be less than or equal to zero")
	}

	return nil
}

// validateInterestRate Returns true if interestRate is greater than zero, else returns false with an error
func validateInterestRate(interestRate float64) error {
	if interestRate <= 0 {
		return fmt.Errorf("value cannot be less than or equal to zero")
	}

	return nil
}

// validateTermLength Returns the total years and months if determined, otherwise returns 0, 0 with an error
func validateTermLength(termLength string) (*TermLength, error) {
	var result *TermLength
	var err error

	split := strings.Split(termLength, " ")

	switch len(split) {
	// One term e.g. [3 years], [12 months]
	case 2:
		term := []string{split[0], split[1]}

		result, err = parseTermLength(term)
		if err != nil {
			return nil, err
		}
	// Two terms e.g. [3 years 9 months], [1 year 12 months]
	case 4:
		terms := [][]string{
			split[:len(split)/2],
			split[len(split)/2:],
		}

		result, err = parseTermLength(terms...)
		if err != nil {
			return nil, err
		}
	// Unrecognised input
	default:
		return nil, fmt.Errorf("could not parse term length")
	}

	// Input was valid but both values entered for years and months was zero
	if result.Years == 0 && result.Months == 0 {
		return nil, fmt.Errorf("years and months cannot both be zero")
	}

	return result, nil
}

// validateInterestPaid returns the PaidAt value if determined, otherwise
// Undefined is returned along with a corresponding error.
func validateInterestPaid(interestPaid string) (PaidAt, error) {
	paidAt := NewPaidAt(interestPaid)
	if paidAt == PaidAt_Undefined {
		return PaidAt_Undefined, fmt.Errorf("input for paid at interval undefined")
	}
	return paidAt, nil
}

// parseTermLength takes in a slice of terms (e.g. [3 years], [[3 years], [2 months]])
// and returns term length. If the terms are not valid or any errors are returned
// during parsing, an error is returned.
func parseTermLength(terms ...[]string) (*TermLength, error) {
	result := new(TermLength)

	var yearFound, monthFound bool

	for _, term := range terms {
		termValue, err := strconv.ParseInt(term[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing value for paid at interval")
		}

		switch v := term[1]; {
		case strings.Contains(v, "year"):
			if yearFound {
				return nil, fmt.Errorf("multiple fields of the same interval detected")
			}
			result.Years = int(termValue)
			yearFound = true
		case strings.Contains(v, "month"):
			if monthFound {
				return nil, fmt.Errorf("multiple fields of the same interval detected")
			}
			result.Months = int(termValue)
			monthFound = true
		default:
			return nil, fmt.Errorf("invalid valid provided for paid at interval")
		}
	}

	return result, nil
}

// calculate performs the calculation to determine the final balance of the term
// deposit given the parameters provided. If an error occurs, a final balance of
// zero is returned along with the error.
func calculate(startDeposit int, interestRate float64, termLength *TermLength, paidAt PaidAt) (float64, error) {
	// Validate PaidAt value is valid
	if paidAt == PaidAt_Undefined {
		return 0, fmt.Errorf("cannot calculate final balance without a defined paid at interval")
	}

	// Make sure TermLength value is valid
	totalMonths := (termLength.Years * 12) + termLength.Months
	if totalMonths <= 0 {
		return 0, fmt.Errorf("total time of investment term cannot be less than 1 month")
	}

	// Paid At Maturity uses a different equation defined in its own function
	if paidAt == PaidAt_AtMaturity {
		return calculateAtMaturity(startDeposit, interestRate, totalMonths), nil
	}

	// Get the divisor for the compound interest equation (custom to the PaidAt interval)
	compoundDivisor := _divisorMap[paidAt]

	finalBalance := float64(startDeposit) * math.Pow(
		1+((interestRate/100)/compoundDivisor),
		compoundDivisor*(float64(totalMonths)/12),
	)

	return finalBalance, nil
}

// calculateAtMaturity performs the calculation for the final balance of a
// maturity term deposit. Uses a different equations to monthly, quarterly,
// and annually hence a separate function.
func calculateAtMaturity(startDeposit int, interestRate float64, totalMonths int) float64 {
	totalInterest := float64(startDeposit) * (interestRate / 100) * (float64(totalMonths) / 12)

	return float64(startDeposit) + totalInterest
}
