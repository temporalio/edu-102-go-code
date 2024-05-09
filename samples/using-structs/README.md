# Sample: Using Structs for Data
This sample demonstrates how to:

* Define structs to represent input and output of an Activity Definition
* Use those structs in the Activity and Workflow Definitions

## Review the Code
This sample introduces an improved version of the translation 
Workflow used in Temporal 101. The translation language can now
be specified as input, so translations are no longer limited to 
Spanish as they were in Temporal 101. The Workflow and Activity
input now include a new parameter, `LanguageCode`, in addition 
to the term being translated. The Workflow output holds the 
translation of the words for Hello and Goodbye in the specified
language. 

The data structures used for input and output in the Workflow and 
Activity are defined in the `shared.go` file. The Workflow and 
Activity Definitions, as well as the code used to start the 
Workflow Execution, have been updated to use these. 

Take a moment to review the code.

## Run the Translation Workflow
When you're ready to run the Workflow, follow the steps below.

1. In one terminal, start the translation microservice by running 
   `go run microservice/translation-service.go`
2. In another terminal, start the Worker by running `go run worker/main.go`
3. In another terminal, execute the Workflow by running 
   `go run start/main.go Pierre fr` (replace `Pierre` with your 
   first name), which should display customized greeting and farewell 
   messages in French.

It's common for a single Workflow Definition to be executed multiple 
times, each time using a different input. Feel free to experiment 
with this by specifying a different language code when starting the 
Workflow Execution. The translation service currently supports the 
following languages:

* `de`: German
* `es`: Spanish
* `fr`: French
* `lv`: Latvian
* `mi`: Maori
* `sk`: Slovak
* `tr`: Turkish
* `zu`: Zulu

