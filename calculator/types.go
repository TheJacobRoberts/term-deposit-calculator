package calculator

import (
	"strings"
)

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

type NormalisedInputValues struct {
	StartDeposit   int
	InterestRate   float64
	TermLength     *TermLength
	PaidAtInterval PaidAt
}

type TermLength struct {
	Years  int
	Months int
}
