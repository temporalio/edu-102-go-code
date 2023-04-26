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

	version := workflow.GetVersion(ctx, "MovedThankYouAfterLoop", workflow.DefaultVersion, 1)
	if version == workflow.DefaultVersion {
		// For workflow executions started before the change, send thank you before the loop
		var notifyConfirmation string
		err = workflow.ExecuteActivity(ctx, SendThankYouToCustomer, input).Get(ctx, &notifyConfirmation)
		if err != nil {
			return "", err
		}
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

		// using 3 seconds instead of 30 days for faster results
		workflow.Sleep(ctx, time.Second*3)
	}

	if version == 1 {
		// for workflow executions started after the change, send thank you after the loop
		var notifyConfirmation string
		err = workflow.ExecuteActivity(ctx, SendThankYouToCustomer, input).Get(ctx, &notifyConfirmation)
		if err != nil {
			return "", err
		}
	}

	result := fmt.Sprintf("Loan for customer '%s' has been fully paid (total=%d)", input.CustomerID, totalPaid)
	return result, nil
}
