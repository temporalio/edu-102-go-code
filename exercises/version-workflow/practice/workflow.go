package loanprocess

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func LoanProcessingWorkflow(ctx workflow.Context, input CustomerInfo) (string, error) {
	logger := workflow.GetLogger(ctx)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 60,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var totalPaid int
	var err error

	// TODO move this code when prompted
	var notifyConfirmation string
	err = workflow.ExecuteActivity(ctx, SendThankYouToCustomer, input).Get(ctx, &notifyConfirmation)
	if err != nil {
		return "", err
	}

	for period := 1; period <= input.NumberOfPeriods; period++ {

		chargeInput := ChargeInput{
			CustomerID:      input.CustomerID,
			Amount:          input.Amount,
			PeriodNumber:    period,
			NumberOfPeriods: input.NumberOfPeriods,
		}

		var chargeConfirmation string
		err = workflow.ExecuteActivity(ctx, ChargeCustomer, chargeInput).Get(ctx, &chargeConfirmation)
		if err != nil {
			return "", err
		}

		totalPaid += chargeInput.Amount
		logger.Info("Payment complete", "Period", period, "Total Paid", totalPaid)

		// TODO change the duration of this Timer when prompted
		workflow.Sleep(ctx, time.Second*3)
	}

	result := fmt.Sprintf("Loan for customer '%s' has been fully paid (total=%d)", input.CustomerID, totalPaid)
	return result, nil
}
