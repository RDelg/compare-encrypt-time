# Choose a minimal base image suitable for Go
FROM golang:1.17-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the dependencies and source files
COPY go.* ./
RUN go mod download

COPY . .

# Build the Go project as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Create a final lightweight alpine image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app

# Execute the binary
ENTRYPOINT ["/app"]
