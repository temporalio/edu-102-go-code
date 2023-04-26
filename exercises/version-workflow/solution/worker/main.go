package main

import (
	"log"
	loanprocess "temporal102/exercises/version-workflow/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, loanprocess.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(loanprocess.LoanProcessingWorkflow)
	w.RegisterActivity(loanprocess.ChargeCustomer)
	w.RegisterActivity(loanprocess.SendThankYouToCustomer)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
