# Exercise 1: Using Structs for Data
During this exercise, you will

* Define structs to represent input and output of an Activity Definition
* Update the Activity and Workflow code to use these structs
* Run the Workflow to ensure that it works as expected

Make your changes to the code in the `practice` subdirectory (look for `TODO` comments that will guide you to where you should make changes to the code). If you need a hint or want to verify your changes, look at the complete version in the `solution` subdirectory.

## Part A: Define the Activity Structs
This exercise provides an improved version of the translation Workflow used in Temporal 101. The Workflow has already been updated to follow the best practice of using structs to represent input parameters and return values. You'll apply what you've learned to do the same for the Activity.

Before continuing with the steps below, take a moment to look at the code in the `shared.go` file to see how the structs are defined for the Workflow. After this, look at the `workflow.go` file to see how these values are passed in and used in the Workflow code. Finally, look at the `start/main.go` to see how the input parameters are created and passed into the Workflow.

Once you're ready to implement something similar for the Activity, continue with the steps below:

1. Edit the `shared.go` file
2. Define a struct called `TranslationActivityInput` to use as an input parameter. 
   1. Define a field named `Term` of type `string`
   2. Define a field named `LanguageCode` of type `string`
3. Define a struct called `TranslationActivityOutput` to use for the result
   1. Define a field named `Translation` of type `string`
4. Save your changes


## Part B: Use the Structs in Your Activity
Now that you have defined the structs, you must update the Activity code to use them.

1. Edit the `activity.go` file
2. Replace the last two input parameters in the `TranslateTerm` function with the struct you defined as input
3. Replace the first output type (string) in the `TranslateTerm` function with the name of the struct you defined as output
4. There are three lines where the function returns the empty string and an error. Replace the empty string in those statements with an empty output struct (i.e., replace `""` with `TranslationActivityOutput{}`) 
5. At the end of the function, create a `TranslationActivityOutput{}` struct and populate its `Translation` field with the `content` variable, which holds the translation returned in the microservice call. 
6. Return the struct created in the previous step
7. Save your changes


## Part C: Update the Workflow Code
You've now updated the Activity code to use the structs. The next step is to update the Workflow code to use these structs where it passes input to the Activity and access its return value.

1. Edit the `workflow.go` file
2. Add a new line to define a `TranslationActivityInput` struct, populating it with the two fields (term and language code) currently passed as input to the first `ExecuteActivity` call
3. Change the variable type used to access the result the first call to `ExecuteActivity` from `string` to `TranslationActivityOutput`
4. Change that `ExecuteActivity` call to use the struct as input instead of the two parameters it now uses
5. Update the `helloMessage` string so that it is based on the `Translation` field from the Activity output struct
6. Repeat the previous four steps for the second call to `ExecuteActivity`, which translates "Goodbye" 
7. Save your changes


## Part D: Run the Translation Workflow
Now that you've made the necessary changes, it's time to run the Workflow to ensure that it works as expected.

1. In one terminal, start the translation microservice by running `go run  microservice/translation-service.go`
2. In another terminal, start the Worker by running `go run worker/main.go`
3. In another terminal, execute the Workflow by running `go run start/main.go Pierre fr` (replace `Pierre` with your first name), which should display customized greeting and farewell messages in French.

If your code didn't work as expected, go back and doublecheck your changes, possibly comparing them to the code in the `solution` directory.

It's common for a single Workflow Definition to be executed multiple times, each time using a different input. Feel free to experiment with this by specifying a different language code when starting the Workflow. The translation service currently supports the following languages:

* `de`: German
* `es`: Spanish
* `fr`: French
* `lv`: Latvian
* `mi`: Maori
* `sk`: Slovak
* `tr`: Turkish
* `zu`: Zulu



### This is the end of the exercise.

