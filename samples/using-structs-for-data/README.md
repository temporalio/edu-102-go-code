# Optional: Using Structs for Data

In this example, you will see how structs are used to represent input and output data of Workflow and Activity Definitions.


## Part A: Observing the Defined Structs in Workflows
This sample provides an improved version of the translation Workflow used in Temporal 101. The Workflow follows the best practice of using structs to represent input parameters and return values. 

Look at the code in the `shared.go` file to see how the structs are defined for the Workflows and Activities. After this, look at the `workflow.go` file to see how these values are passed in and used in the Workflow code. Finally, look at the `start/main.go` to see how the input parameters are created and passed into the Workflow.


## Part B: Observing the Defined Structs in Activities
Now let's take a look at how we used structs to represent input and output data in Activity definitions. 

Take a look at the `activity.go` file to see how the `TranslateTerm` function takes in the `TranslationActivityInput` struct as an input parameter. Also notice how that function returns a `TranslationActivityOutput` struct for the output.


## Part C: Run the Translation Workflow
To run the Workflow:

1. In one terminal, start the translation microservice by running `go run  microservice/translation-service.go`
2. In another terminal, start the Worker by running `go run worker/main.go`
3. In another terminal, execute the Workflow by running `go run start/main.go Pierre fr` (replace `Pierre` with your first name), which should display customized greeting and farewell messages in French.

It's common for a single Workflow Definition to be executed multiple times, each time using a different input. Feel free to experiment with this by specifying a different language code when starting the Workflow. The translation service currently supports the following languages:

* `de`: German
* `es`: Spanish
* `fr`: French
* `lv`: Latvian
* `mi`: Maori
* `sk`: Slovak
* `tr`: Turkish
* `zu`: Zulu



### This is the end of the sample.

