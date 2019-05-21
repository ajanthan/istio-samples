FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
# Create appuser.
RUN adduser -D -g '' appuser
ENV GO111MODULE on
WORKDIR $GOPATH/src/github.com/ajanthan/istio-samples
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Using go mod.
# RUN go mod download
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build  -ldflags="-w -s" -o /go/bin/echo-service services/echo/cmd/echo.go

FROM golang:alpine

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /go/bin/echo-service /go/bin/echo-service

# Use an unprivileged user.
USER appuser

# Expose service
EXPOSE 8080

# Run the hello binary.
ENTRYPOINT ["/go/bin/echo-service"]
