FROM golang:1.18.2-bullseye

# RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.

COPY . .

# Build the Go app
# RUN apt install git
RUN go build -o ./bin/app .
# RUN go mod vendor


# This container exposes port 8080 to the outside world
EXPOSE 12345

# Run the binary program produced by `go install`
CMD ["./bin/app"]

