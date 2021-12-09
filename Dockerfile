FROM golang:1.17-alpine AS dep

# Add the module files and download dependencies.
COPY ./go.mod /go/src/app/go.mod
COPY ./go.sum /go/src/app/go.sum
WORKDIR /go/src/app

RUN go mod download
# Add the shared packages.
# COPY ./data /go/src/app/data
# COPY ./util /go/src/app/util
