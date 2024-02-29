# Start from the official Golang image
FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY auth /app/auth
COPY config /app/config
COPY db /app/db
COPY handler /app/handler
COPY memory /app/memory
COPY middleware /app/middleware
COPY models /app/models
COPY rate_limit /app/rate_limit
COPY repository /app/repository
COPY router /app/router
COPY service /app/service
COPY vendor /app/vendor
COPY main.go /app/main.go


# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
