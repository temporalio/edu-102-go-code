# Age Estimation Workflow
This example provides a Workflow that can estimate someone's age, 
based on their given name. It uses a single Activity implementation, 
which calls a remote API to retrieve this estimation.

In addition to the Worker, it also includes a file `start/main.go` 
that you can use to run the Workflow (pass a name to use as input 
on the command line when invoking it). 

```
$ go run start/main.go Betty
```

This will output a message with the name and estimated age:

```
Betty has an estimated age of 76
```

Additionally, this example provides tests for the Workflow 
and Activity code.
