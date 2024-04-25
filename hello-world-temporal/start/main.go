package main

import (
	"context"
	"fmt"
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	// create a client and dial in to Temporal Cluster
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	// start the workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete workflow", err)
	}

	// get the result
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}

	printResult(greeting, we.GetID(), we.GetRunID())
}

func printResult(greeting, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
