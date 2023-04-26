package loanprocess

const TaskQueueName = "loan-processing-workflow-taskqueue"

type ChargeInput struct {
	CustomerID      string
	Amount          int
	PeriodNumber    int
	NumberOfPeriods int
}
