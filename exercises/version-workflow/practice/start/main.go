package main

import (
	"context"
	"fmt"
	"log"
	"os"
	loanprocess "temporal102/exercises/version-workflow/practice"

	"go.temporal.io/sdk/client"
)

func main() {
	// look up the customer information using the ID specified as a command-line argument
	if len(os.Args) <= 1 {
		log.Fatalln("Must specify customer ID as command-line argument")
	}
	customerID := os.Args[1]

	db := loanprocess.CustomerInfoDB()
	info, err := db.Get(customerID)
	if err != nil {
		msg := fmt.Sprintf("Error looking up customer ID %s, reason: %v", customerID, err)
		log.Fatalln(msg)
	}

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "loan-processing-workflow-customer-" + info.CustomerID,
		TaskQueue: loanprocess.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, loanprocess.LoanProcessingWorkflow, info)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
