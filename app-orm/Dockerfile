FROM dependencies AS builder

# Copy the application source code.
COPY . /go/src/app/.

# RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0
RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# Build the application.
# RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 \
# RUN go build -o /go/bin/app-orm /go/src/app/.
# ENTRYPOINT [ "/go/bin/web-server" ]
 ENTRYPOINT ["/root/wait_for_db.sh", "ash" ]


FROM alpine:latest

COPY wait_for_db.sh /root/wait_for_db.sh
# COPY create_database.sh /root/create_database.sh

RUN apk update
RUN apk add mysql-client
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
RUN ln -s /usr/local/bin/migrate /bin/migrate
# COPY --from=builder /go/bin/app-orm /bin/app-orm
#ENTRYPOINT ["/root/wait_for_db.sh", "/root/create_database.sh" ]
# ENTRYPOINT ["/root/wait_for_db.sh", "ash" ]
