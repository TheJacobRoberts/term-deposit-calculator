# Term-Deposit-Calculator

A simple term deposit calculator.


## How to Run?

To run the application, simply clone the project, go to the project directory and run:

```go run main.go -start-deposit=startDeposit -interest-rate=interestRate -term-length=termLength -interest-paid=interestPaid```

## Flags

- start-deposit int: The starting amount for the term deposit e.g. 10000 ($10,000).
- interest-rate float64: The interest rate for the term deposit e.g. 1.1 (1.1%).
- term-length string: The length of the term deposit. Can provide both years and months e.g. "3 years", "1 month", "4 years 6 months".
- interest-paid string: The rate at which interest is paid at for the term deposit. Valid values are "monthly", "quarterly" "annually", or "maturity".