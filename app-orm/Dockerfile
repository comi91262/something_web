FROM dependencies AS builder

# Copy the application source code.
COPY . /go/src/app/.

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/app-orm /go/src/app/.

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/healthcheck /go/src/app/healthcheck/main.go

ENTRYPOINT [ "/go/bin/app-orm" ]

FROM alpine:latest

RUN apk add --no-cache mysql-client

WORKDIR /var/www

COPY --from=builder /go/src/app/static /var/www/static
COPY --from=builder /go/bin/app-orm /var/www/app-orm

COPY --from=builder /go/bin/healthcheck /bin/healthcheck
COPY --from=builder /go/src/app/wait_for_db.sh /bin/wait_for_db.sh

ENTRYPOINT ["/bin/healthcheck", "/var/www/app-orm" ]
