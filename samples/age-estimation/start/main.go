package main

import (
	"context"
	"log"
	"os"
	example "temporal102/samples/age-estimation"

	"go.temporal.io/sdk/client"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Must specify a name as the command-line argument")
	}
	name := os.Args[1]

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "testing-estimate-age-example",
		TaskQueue: example.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, example.EstimateAge, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	log.Println("Workflow result:", result)
}
