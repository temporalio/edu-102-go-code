package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulCompleteFrenchTranslation(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}

	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(TranslateTerm)

	workflowInput := TranslationWorkflowInput{
		Name:         "Pierre",
		LanguageCode: "fr",
	}

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	assert.True(t, env.IsWorkflowCompleted())

	var result TranslationWorkflowOutput
	env.GetWorkflowResult(&result)

	assert.Equal(t, "Bonjour, Pierre", result.HelloMessage)
	assert.Equal(t, "Au revoir, Pierre", result.GoodbyeMessage)
}
