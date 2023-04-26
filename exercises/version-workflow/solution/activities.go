package loanprocess

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func ChargeCustomer(ctx context.Context, input ChargeInput) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("*** Charging customer ***",
		"CustomerID:", input.CustomerID,
		"Amount:", input.Amount,
		"PeriodNumber:", input.PeriodNumber,
		"NumberOfPeriods:", input.NumberOfPeriods)

	// just pretend that we charged them
	confirmation := fmt.Sprintf("Charged $%d to customer '%s'", input.Amount, input.CustomerID)

	return confirmation, nil
}

func SendThankYouToCustomer(ctx context.Context, input CustomerInfo) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("*** Sending thank you message To Customer ***",
		"CustomerID:", input.CustomerID,
		"EmailAddress:", input.EmailAddress)

	// just pretend that we e-mailed them
	confirmation := fmt.Sprintf("Sent thank you message to customer '%s'", input.CustomerID)

	return confirmation, nil
}
