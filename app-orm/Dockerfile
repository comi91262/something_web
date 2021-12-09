FROM dependencies AS builder

# Copy the application source code.
COPY . /go/src/app/.

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/app-orm /go/src/app/.
ENTRYPOINT [ "/go/bin/web-server" ]

FROM alpine:latest
COPY wait_for_db.sh /root/wait_for_db.sh

RUN apk update
RUN apk add mysql-client

COPY --from=builder /go/bin/app-orm /bin/app-orm
ENTRYPOINT ["/root/wait_for_db.sh", "/bin/app-orm" ]
