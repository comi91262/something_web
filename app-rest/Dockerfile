FROM dependencies AS builder

# Copy the application source code.
COPY . /go/src/app/.

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/app-rest /go/src/app/.

ENTRYPOINT [ "/go/bin/app-rest" ]

FROM alpine:latest

RUN apk add --no-cache mysql-client

WORKDIR /var/www

COPY --from=builder /go/bin/app-rest /var/www/app-rest
COPY --from=builder /go/src/app/wait_for_db.sh /bin/wait_for_db.sh

ENTRYPOINT ["/bin/wait_for_db.sh", "/var/www/app-rest" ]
