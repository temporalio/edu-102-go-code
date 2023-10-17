# Exercise 2: Testing the Translation Workflow
During this exercise, you will

* Run a unit test provided for the `TranslateTerm` Activity
* Develop and run your own unit test for the `TranslateTerm` Activity
* Write assertions for a Workflow test 
* Uncover, diagnose, and fix a bug in the Workflow Definition
* Observe the time-skipping feature in the Workflow test environment

Make your changes to the code in the `practice` subdirectory (look for 
`TODO` comments that will guide you to where you should make changes to 
the code). If you need a hint or want to verify your changes, look at 
the complete version in the `solution` subdirectory.

## Part A: Running a Test

We have provided a unit test for the `TranslateTerm` Activity
to get you started. This test verifies that the Activity correctly 
translates the term "Hello" to German. Take a moment to study the 
test, which you'll find in the `activity_test.go` file. Since the 
test runs the Activity, which in turn calls the microservice to do 
the translation, you'll begin by starting that.

1. Open a new terminal and run `go run microservice/translation-service.go` 
2. Run the `go test` command to execute the provided test

## Part B: Write and Run Another Test for the Activity

Now it's time to develop and run your own unit test, this time 
verifying that the Activity correctly supports the translation 
of a different word in a different language.

1. Edit the `activity_test.go` file
2. Copy the `TestSuccessfulTranslateActivityHelloGerman` function, 
   renaming the new function as `TestSuccessfulTranslateActivityGoodbyeLatvian`
3. Change the term for the input from `Hello` to `Goodbye` 
4. Change the language code for the input from `de` (German) to `lv` (Latvian)
5. Assert that translation returned by the Activity is `Ardievu` 

## Part C: Test the Activity with Invalid Input

In addition to verifying that your code behaves correctly when used as 
you intended, it is sometimes also helpful to verify its behavior with 
unexpected input. The example below does this, testing that the Activity 
returns the appropriate error when called with an invalid language code. 

```go
func TestFailedTranslateActivityBadLanguageCode(t *testing.T) {
	testSuite := testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(TranslateTerm)

	input := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: "xq",
	}

	_, err := env.ExecuteActivity(TranslateTerm, input)

	// Assert that execution returned an error
	assert.Error(t, err)

	// Assert that this is a Temporal Application Error
	var applicationErr *temporal.ApplicationError
	assert.True(t, errors.As(err, &applicationErr))

	// Assert that the error has the expected message, which identifies
	// the invalid language code as the cause
	assert.Equal(t, "HTTP Error 400: Unknown language code 'xq'\n", applicationErr.Message())
}
```

Take a moment to study this code, and then continue with the 
following steps:

1. Edit the `activity_test.go` file
2. Uncomment the imports on lines 4 and 8
3. Copy the entire `TestFailedTranslateActivityBadLanguageCode` function
   provided above and paste it at the bottom of the `activity_test.go` file 
4. Save the changes
5. Run `go test` again to run this new test, in addition to the others


## Part D: Test a Workflow Definition

1. Edit the `workflow_test.go` file
2. Uncomment the import on line 6
3. Remove `IGNORE` from the function name on line 10 (this was added 
   so the test didn't run in the earlier parts of this exercise, since 
   the function name must begin with `Test` to be recognized as a test)
4. Add assertions for the following conditions
   * The Workflow Execution has completed
   * The `HelloMessage` field in the result is `Bonjour, Pierre`
   * The `GoodbyeMessage` field in the result is `Au revoir, Pierre`
5. Save your changes
6. Run `go test`. This will fail, due to a bug in the Workflow Definition.
7. Find and fix the bug in the Workflow Definition
8. Run the `go test` command again to verify that you fixed the bug

There are two things to note about this test.

First, the test completes in under a second, even though the Workflow 
Definition contains a `workflow.Sleep` call that adds a 15-second delay 
to the Workflow Execution. This is because of the time-skipping feature
provided by the test environment.

Second, calls to `RegisterActivity` near the top of the test indicate 
that the Activity Definitions are executed as part of this Workflow 
test. As you learned, you can test your Workflow Definition in isolation 
from the Activity implementations by using mocks. The optional exercise 
that follows provides an opportunity to try this for yourself.


### This is the end of the exercise.


## (Optional) Using Mock Activities in a Workflow Test

If you have time and would like an additional challenge, 
continue with the following steps.

1. Make a copy of the existing Workflow Test by running 
   `cp workflow_test.go workflow_mock_test.go`
2. Edit the `workflow_mock_test.go` file
3. Add an import for `"github.com/stretchr/testify/mock"`
4. Rename the test function to `TestSuccessfulTranslationWithMocks`
5. Delete the line used to register the Activity. 
   This is unnecessary in a Workflow test that uses mock
   objects for the Activity, since the *actual* Activity 
   Definition is never executed.
6. Make the following changes between where the struct representing
   workflow input is defined and where `env.ExecuteWorkflow` is called
   * Create and populate an instance of the `TranslationActivityInput`
     struct to represent input passed to the Activity when translating 
     the greeting
   * Create and populate an instance of the `TranslationActivityOutput`
     struct to represent output returned by the Activity when translating 
     the greeting
   * Create a mock that represents the `TranslateTerm` Activity, 
     which will return the output struct you created when called 
     with the input struct you created
   * Repeat the above three steps, this time creating structs and 
     a mock for the goodbye message
7. Save your changes
8. Run `go test` to run the tests
