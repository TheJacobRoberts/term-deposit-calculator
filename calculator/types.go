package calculator

import (
	"strings"
)

// PaidAt is an 'enum' which can be of the following values:
// Undefined, Monthly, Quarterly, Annually, or AtMaturity.
// It represents when the interval for when interest is paid.
type PaidAt int

const (
	PaidAt_Undefined PaidAt = iota
	PaidAt_Monthly
	PaidAt_Quarterly
	PaidAt_Annually
	PaidAt_AtMaturity
)

func NewPaidAt(value string) PaidAt {
	switch v := strings.ToLower(value); {
	case strings.Contains(v, "monthly"):
		return PaidAt_Monthly
	case strings.Contains(v, "quarterly"):
		return PaidAt_Quarterly
	case strings.Contains(v, "annually"):
		return PaidAt_Annually
	case strings.Contains(v, "maturity"):
		return PaidAt_AtMaturity
	}
	return PaidAt_Undefined
}

func (v PaidAt) String() string {
	switch v {
	case PaidAt_Monthly:
		return "monthly"
	case PaidAt_Quarterly:
		return "quarterly"
	case PaidAt_Annually:
		return "annually"
	case PaidAt_AtMaturity:
		return "at maturity"
	}
	return "undefined"
}

// NormalisedInputValues defines the normalised values of user inputs used in
// the calculator.
type NormalisedInputValues struct {
	// The starting amount of the investment e.g. 10000 ($10,000)
	StartDeposit int
	// The interest rate pct. of the investment e.g. 1.1 (1.1%)
	InterestRate float64
	// The term length of investment e.g. 3 years, 2 months
	TermLength *TermLength
	// The interval which the interest is paid at on the investment e.g. quarterly
	PaidAtInterval PaidAt
}

// TermLength defines the amount of years and months of a term
type TermLength struct {
	Years  int
	Months int
}
