FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/test-task

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/test-task-app advanced/cmd/test-task

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/test-task/out/test-task-app /app/test-task-app

# This container exposes port 8080 to the outside world
EXPOSE 1323

# Run the binary program produced by `go install`
CMD ["/app/test-task-app"]