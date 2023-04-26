package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("SayHelloGoodbye Workflow Invoked", "Name", input.Name)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 120,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	logger.Debug("Preparing to translate Hello", "LanguageCode", input.LanguageCode)
	helloInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: input.LanguageCode,
	}
	var helloResult TranslationActivityOutput
	err := workflow.ExecuteActivity(ctx, TranslateTerm, helloInput).Get(ctx, &helloResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	helloMessage := fmt.Sprintf("%s, %s", helloResult.Translation, input.Name)

	// Wait a little while before saying goodbye
	logger.Debug("Sleeping between translation calls")
	workflow.Sleep(ctx, time.Second*15)

	logger.Debug("Preparing to translate Goodbye", "LanguageCode", input.LanguageCode)
	goodbyeInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: input.LanguageCode,
	}
	var goodbyeResult TranslationActivityOutput
	err = workflow.ExecuteActivity(ctx, TranslateTerm, goodbyeInput).Get(ctx, &goodbyeResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	goodbyeMessage := fmt.Sprintf("%s; %s", goodbyeResult.Translation, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}
