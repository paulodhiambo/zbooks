# Build stage
FROM golang:1.21 AS builder

WORKDIR /app

COPY . ./
RUN go mod download

# Change directory to ./cmd/userservice
#WORKDIR /app/cmd/userservice

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /app/cmd/userservice

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=builder /app/cmd/userservice/app .

CMD ["./app"]
