FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

# Set the working directory for the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on (if applicable)
# EXPOSE 8080

# Define the command to run when the container starts
CMD ["./main"]
