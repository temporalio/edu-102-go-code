package pizza

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestGetDistanceTwoLineAddress(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(GetDistance)

	input := Address{
		Line1:      "701 Mission Street",
		Line2:      "Apartment 9C",
		City:       "San Francisco",
		State:      "CA",
		PostalCode: "94103",
	}

	future, err := env.ExecuteActivity(GetDistance, input)
	require.NoError(t, err)

	var output Distance
	assert.NoError(t, future.Get(&output))

	assert.Equal(t, 20, output.Kilometers)
}

func TestGetDistanceOneLineAddress(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(GetDistance)

	input := Address{
		Line1:      "917 Delores Street",
		City:       "San Francisco",
		State:      "CA",
		PostalCode: "94103",
	}

	future, err := env.ExecuteActivity(GetDistance, input)
	require.NoError(t, err)

	var output Distance
	assert.NoError(t, future.Get(&output))

	assert.Equal(t, 8, output.Kilometers)
}

func TestSendBillTypicalOrder(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(SendBill)

	input := Bill{
		CustomerID:  12983,
		OrderNumber: "PI314",
		Description: "2 large cheese pizzas",
		Amount:      2600, // amount does not qualify for discount
	}

	future, err := env.ExecuteActivity(SendBill, input)
	require.NoError(t, err)

	var confirmation OrderConfirmation
	assert.NoError(t, future.Get(&confirmation))

	assert.Equal(t, "PI314", confirmation.OrderNumber)
	assert.Equal(t, 2600, confirmation.Amount)
}

func TestSendBillFailsWithNegativeAmount(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(SendBill)

	input := Bill{
		CustomerID:  21974,
		OrderNumber: "OU812",
		Description: "1 large sausage pizza",
		Amount:      -1000,
	}

	_, err := env.ExecuteActivity(SendBill, input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid charge amount: -1000")
}
