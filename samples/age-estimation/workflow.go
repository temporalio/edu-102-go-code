package example

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func EstimateAge(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var age int
	err := workflow.ExecuteActivity(ctx, RetrieveEstimate, name).Get(ctx, &age)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s has an estimated age of %d", name, age)

	return result, nil
}
