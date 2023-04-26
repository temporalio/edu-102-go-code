package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulTranslationWithMocks(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}

	env := s.NewTestWorkflowEnvironment()

	workflowInput := TranslationWorkflowInput{
		Name:         "Pierre",
		LanguageCode: "fr",
	}

	helloInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: workflowInput.LanguageCode,
	}
	helloOutput := TranslationActivityOutput{
		Translation: "Bonjour",
	}
	env.OnActivity(TranslateTerm, mock.Anything, helloInput).Return(helloOutput, nil)

	goodbyeInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: workflowInput.LanguageCode,
	}
	goodbyeOutput := TranslationActivityOutput{
		Translation: "Au revoir",
	}
	env.OnActivity(TranslateTerm, mock.Anything, goodbyeInput).Return(goodbyeOutput, nil)

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	assert.True(t, env.IsWorkflowCompleted())

	var result TranslationWorkflowOutput
	env.GetWorkflowResult(&result)

	assert.Equal(t, "Bonjour, Pierre", result.HelloMessage)
	assert.Equal(t, "Au revoir, Pierre", result.GoodbyeMessage)
}
