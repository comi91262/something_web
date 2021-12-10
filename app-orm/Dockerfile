FROM dependencies AS builder

# Copy the application source code.
COPY . /go/src/app/.

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/app-orm /go/src/app/.

ENTRYPOINT [ "/go/bin/app-orm" ]

FROM alpine:latest


RUN apk add --no-cache mysql-client
COPY wait_for_db.sh /bin/wait_for_db.sh
COPY --from=builder /go/bin/app-orm /bin/app-orm


ENTRYPOINT ["/bin/wait_for_db.sh", "/bin/app-orm" ]
