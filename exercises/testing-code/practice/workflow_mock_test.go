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

	// Mock hello activity
	mockHelloActivityInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: "fr",
	}
	mockHelloActivityOutput := TranslationActivityOutput{
		Translation: "MockHelloActivityOutput",
	}
	env.OnActivity(TranslateTerm, mock.Anything, mockHelloActivityInput).Return(mockHelloActivityOutput, nil)

	// Mock goodbye activity
	mockGoodbyeActivityInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: "fr",
	}
	mockGoodbyeActivityOutput := TranslationActivityOutput{
		Translation: "MockGoodbyeActivityOutput",
	}
	env.OnActivity(TranslateTerm, mock.Anything, mockGoodbyeActivityInput).Return(mockGoodbyeActivityOutput, nil)

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	// Assert that Workflow Execution completed
	assert.True(t, env.IsWorkflowCompleted())

	var result TranslationWorkflowOutput
	env.GetWorkflowResult(&result)

	// Assert that the HelloMessage field in the
	assert.Equal(t, "MockHelloActivityOutput, Pierre", result.HelloMessage)

	// Assert that the GoodbyeMessage field in the
	assert.Equal(t, "MockGoodbyeActivityOutput, Pierre", result.GoodbyeMessage)
}
