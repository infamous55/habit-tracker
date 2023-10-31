# Build the application from source
FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/server

# Deploy the aplication binary
FROM alpine:latest

COPY --from=build /go/bin/app /app

CMD ["/app"]
