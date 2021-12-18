FROM golang:1.17 AS dep

# Add the module files and download dependencies.
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum

WORKDIR /go/src/app

RUN go mod download

# Add the shared packages.
COPY ./tools /go/src/app/tools

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/healthcheck /go/src/app/tools/healthcheck/main.go
