package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	translation "temporal102/exercises/testing-code/practice"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "translation-workflow",
		TaskQueue: translation.TaskQueueName,
	}

	if len(os.Args) <= 2 {
		log.Fatalln("Must specify name and language code as command-line arguments")
	}

	input := translation.TranslationWorkflowInput{
		Name:         os.Args[1],
		LanguageCode: os.Args[2],
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, translation.SayHelloGoodbye, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result translation.TranslationWorkflowOutput
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln("Unable to format result in JSON format", err)
	}
	log.Printf("Workflow result: %s\n", string(data))
}
