package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func Test_EstimateAge_WithMockActivity(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// This mock represents the RetrieveEstimate Activity, and if 
	// passed "Betty" as input, it will return 76 as output.
	env.OnActivity(RetrieveEstimate, mock.Anything, "Betty").Return(76, nil)

	env.ExecuteWorkflow(EstimateAge, "Betty")
	assert.True(t, env.IsWorkflowCompleted())

	var result string
	assert.NoError(t, env.GetWorkflowResult(&result))
	expected := "Betty has an estimated age of 76"
	assert.Equal(t, expected, result)
}
