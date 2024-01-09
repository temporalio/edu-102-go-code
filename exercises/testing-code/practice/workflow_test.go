package translation

import (
	"testing"

	//	"github.com/stretchr/testify/assert"
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

	// Assert that Workflow Execution completed
	assert.True(t, env.IsWorkflowCompleted())

	var result TranslationWorkflowOutput
	env.GetWorkflowResult(&result)

	// TODO: Assert that the HelloMessage field in the
	//       result is: Bonjour, Pierre
	assert.Equal(t, "Bonjour, Pierre", result.HelloMessage)

	// TODO: Assert that the GoodbyeMessage field in the
	//       result is: Au revoir, Pierre
	assert.Equal(t, "Au revoir, Pierre", result.GoodbyeMessage)
}
