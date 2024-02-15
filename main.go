package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/thejacobroberts/term-deposit-calculator/calculator"
)

var (
	startDeposit int
	interestRate float64
	termLength   string
	interestPaid string
)

func init() {
	flag.IntVar(&startDeposit, "start-deposit", 0, "the starting balance for the term deposit (e.g. 10,000)")
	flag.Float64Var(&interestRate, "interest-rate", 0, "the interest rate for the term deposit (e.g. 1.1)")
	flag.StringVar(&termLength, "term-length", "", "the term length for the term deposit (e.g. 3 years)")
	flag.StringVar(&interestPaid, "interest-paid", "", "when the interest is paid on the term deposit (monthly, quarterly, annually, at maturity)")
}

func main() {
	flag.Parse()

	fmt.Println("[ Term Deposit Calculator ]")

	fmt.Println("Users inputs:")
	fmt.Print(
		fmt.Sprintf("\tstart-deposit: '%v'\n", startDeposit),
		fmt.Sprintf("\tinterest-rate: '%v'\n", interestRate),
		fmt.Sprintf("\tterm-length: '%v'\n", termLength),
		fmt.Sprintf("\tinterest-paid: '%v'\n", interestPaid),
	)

	normalisedInputs, errs := calculator.Validate(startDeposit, interestRate, termLength, interestPaid)
	if len(errs) > 0 {
		log.Fatalf("failed to validate input values - received the following errors:\n%v", errors.Join(errs...)) // TODO make this look prettier
	}

	fmt.Println("Normalised inputs:")
	fmt.Print(
		fmt.Sprintf("\tStart Deposit: $%d\n", startDeposit),
		fmt.Sprintf("\tInterest Rate: %.2f%%\n", interestRate),
		fmt.Sprintf("\tTerm Length: %s\n", termLength),
		fmt.Sprintf("\tPaid: %s\n", normalisedInputs.PaidAtInterval.String()),
	)

	finalBalance, err := calculator.Calculate(
		normalisedInputs.StartDeposit,
		normalisedInputs.InterestRate,
		normalisedInputs.TermLength,
		normalisedInputs.PaidAtInterval,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Final Balance:\n\t$%.2f\n", finalBalance)
}
