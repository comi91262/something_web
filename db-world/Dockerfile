FROM dependencies AS builder

COPY migrate.sh /go/src/app/.
COPY migrations /go/src/app/migrations

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0

ENTRYPOINT ["/go/bin/healthcheck", "./migrate.sh"]
