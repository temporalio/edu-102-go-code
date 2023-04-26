package translation

import (
	"testing"

	//	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func IGNORETestSuccessfulCompleteFrenchTranslation(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}

	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(SayHelloGoodbye)
	env.RegisterActivity(TranslateTerm)

	workflowInput := TranslationWorkflowInput{
		Name:         "Pierre",
		LanguageCode: "fr",
	}

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	// TODO: Assert that Workflow Execution completed

	var result TranslationWorkflowOutput
	env.GetWorkflowResult(&result)

	// TODO: Assert that the HelloMessage field in the
	//       result is: Bonjour, Pierre

	// TODO: Assert that the GoodbyeMessage field in the
	//       result is: Au revoir, Pierre
}
