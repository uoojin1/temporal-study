package main

import (
	"hello-world-temporal/app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// this worker hosts the workflow and activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.GreetingWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	// start listening to the task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
