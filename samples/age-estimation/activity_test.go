package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

func Test_RetrieveEstimateBob(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(RetrieveEstimate)

	val, err := env.ExecuteActivity(RetrieveEstimate, "Bob")
	assert.NoError(t, err)

	var result int
	assert.NoError(t, val.Get(&result))
	assert.Equal(t, 70, result)
}
