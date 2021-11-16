FROM golang:1.17.3-alpine AS builder

# Create a working directory on root
WORKDIR /build

# Coping everything from current directory to working directory
COPY . .

# Build binary for the service
RUN go build -mod=vendor -o ./user-service ./cmd/app.go


FROM golang:1.17.3-alpine
WORKDIR /app
COPY --from=builder /build/user-service ./user-service

# Run service when container about to start
ENTRYPOINT [ "./user-service" ]