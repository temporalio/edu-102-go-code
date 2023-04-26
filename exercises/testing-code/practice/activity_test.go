package translation

import (
	//	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	//	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulTranslateActivityHelloGerman(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(TranslateTerm)

	input := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: "de",
	}

	future, err := env.ExecuteActivity(TranslateTerm, input)
	assert.NoError(t, err)

	var output TranslationActivityOutput
	assert.NoError(t, future.Get(&output))
	assert.Equal(t, "Hallo", output.Translation)
}
