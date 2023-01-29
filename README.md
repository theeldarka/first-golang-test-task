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
docker run -p 8090:1323 theeldarka/first-golang-test-task:latest
```

Go to [http://localhost:8090](http://localhost:8090n)

Let's check the middleware:
```shell
curl http://localhost:8090 --header User-Role:\ Super-Admin
```
Then, look at the console logs
```json
{"time":"2023-01-29T21:38:59.080717517Z","level":"WARN","prefix":"-","file":"middleware.go","line":"14","message":"red button user detected"}
```