package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

// This test case uses the actual Activity implementation,
// rather than a mock object, so it actually calls the API
// when tested.
func Test_EstimateAge_EndToEnd(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.RegisterActivity(RetrieveEstimate)

	env.ExecuteWorkflow(EstimateAge, "Betty")
	assert.True(t, env.IsWorkflowCompleted())

	var result string
	assert.NoError(t, env.GetWorkflowResult(&result))
	expected := "Betty has an estimated age of 76"
	assert.Equal(t, expected, result)
}
