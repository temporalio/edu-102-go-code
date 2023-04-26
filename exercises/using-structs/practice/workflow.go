package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 45,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// TODO Create your Activity input struct and populate it with the last
	//      two fields from the ExecuteActivity call below

	// TODO Replace "string" below with your Activity output struct type
	var helloResult string

	// TODO Use your input struct in the ExecuteActivity call below
	err := workflow.ExecuteActivity(ctx, TranslateTerm, "Hello", input.LanguageCode).Get(ctx, &helloResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	// TODO Update the middle parameter to use the Translation field from the Activity output struct
	helloMessage := fmt.Sprintf("%s, %s", helloResult, input.Name)

	// TODO Create your Activity input struct and populate it with the last
	//      two fields from the ExecuteActivity call below

	// TODO Replace "string" below with your Activity output struct type
	var goodbyeResult string

	// TODO Use your input struct in the ExecuteActivity call below
	err = workflow.ExecuteActivity(ctx, TranslateTerm, "Goodbye", input.LanguageCode).Get(ctx, &goodbyeResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	// TODO Update the middle parameter to use the Translation field from the Activity output struct
	goodbyeMessage := fmt.Sprintf("%s, %s", goodbyeResult, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}
