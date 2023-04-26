package main

import (
	"log"
	translation "temporal102/exercises/using-structs/practice"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, translation.TaskQueueName, worker.Options{})

	w.RegisterWorkflow(translation.SayHelloGoodbye)
	w.RegisterActivity(translation.TranslateTerm)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
