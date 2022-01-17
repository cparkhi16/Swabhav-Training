# Start from golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on
EXPOSE 9000
# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./


RUN go mod download


COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app .


# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file
COPY --from=builder /app/bin/app .

# Run executable
CMD ["./app"]
