# Test Task in Go

You have 1 hour to solve the task. 
If you are lucky enough and have time left, please try to optimise your solution or highlight some edge cases.

### Description

1. You should create a handler which sends how many days left **until 1 Jan 2025** and response with **HTTP 200 OK** status code.
2. Build a middleware, which checks HTTP header `User-Role` presents and contains `admin` and prints `red button user detected` to the console (using default log package or any 3rd party) if so.

Think about coding standards and best practices. 

In general, the task is about providing a working solution, but if you have time to create a good code-structure, we'd love it!

Thanks, [@spatecon](https://github.com/spatecon) for the task!

## Run solution
```shell
docker run theeldarka/first-golang-test-task:latest
```