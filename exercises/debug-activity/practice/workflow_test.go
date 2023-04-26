package pizza

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulPizzaOrder(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}
	env := s.NewTestWorkflowEnvironment()

	order := *createPizzaOrderForTest()

	// For this test, any address will have a distance of 10 kilometer, which
	// is within the delivery area
	distance := Distance{
		Kilometers: 10,
	}
	env.OnActivity(GetDistance, mock.Anything, mock.Anything).Return(distance, nil)

	confirmation := OrderConfirmation{
		OrderNumber:        order.OrderNumber,
		ConfirmationNumber: "AB9923",
		Status:             "SUCCESS",
		BillingTimestamp:   time.Now().Unix(),
		Amount:             2500,
	}
	env.OnActivity(SendBill, mock.Anything, mock.Anything).Return(confirmation, nil)

	env.ExecuteWorkflow(PizzaWorkflow, order)

	assert.True(t, env.IsWorkflowCompleted())

	var result OrderConfirmation
	env.GetWorkflowResult(&result)

	assert.Equal(t, "Z1238", result.OrderNumber)
	assert.Equal(t, "SUCCESS", result.Status)
	assert.Equal(t, "AB9923", result.ConfirmationNumber)
	assert.Equal(t, 2500, result.Amount)
	assert.NotEmpty(t, result.BillingTimestamp)
}

func TestFailedPizzaOrderCustomerOutsideDeliveryArea(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}
	env := s.NewTestWorkflowEnvironment()

	order := *createPizzaOrderForTest()

	distance := Distance{
		Kilometers: 30, // too far away
	}
	env.OnActivity(GetDistance, mock.Anything, mock.Anything).Return(distance, nil)

	// NOTE there is no Mock for the SendBill Activity because it won't be
	// called, given that the Workflow returns an error due to the distance.
	env.ExecuteWorkflow(PizzaWorkflow, order)

	assert.True(t, env.IsWorkflowCompleted())

	// Since the Workflow failed, trying to access its result fails
	var result OrderConfirmation
	assert.Error(t, env.GetWorkflowResult(&result))

	// When the Workflow returns an error during its execution, Temporal
	// wraps it in a Temporal-specific WorkflowExecutionError type, so we
	// must unwrap this to retrieve the error returned in the Workflow code.
	workflowErr := env.GetWorkflowError()
	assert.EqualError(t, errors.Unwrap(workflowErr), "Customer lives too far away for delivery")
}

func createPizzaOrderForTest() *PizzaOrder {
	customer := Customer{
		CustomerID: 12983,
		Name:       "María García",
		Email:      "maria1985@example.com",
		Phone:      "415-555-7418",
	}

	address := Address{
		Line1:      "701 Mission Street",
		Line2:      "Apartment 9C",
		City:       "San Francisco",
		State:      "CA",
		PostalCode: "94103",
	}

	p1 := Pizza{
		Description: "Large, with pepperoni",
		Price:       1500,
	}

	p2 := Pizza{
		Description: "Small, with mushrooms and onions",
		Price:       1000,
	}

	items := []Pizza{p1, p2}

	order := PizzaOrder{
		OrderNumber: "Z1238",
		Customer:    customer,
		Items:       items,
		Address:     address,
		IsDelivery:  true,
	}

	return &order
}
