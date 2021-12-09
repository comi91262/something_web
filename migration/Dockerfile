FROM dependencies AS builder

COPY wait_for_db.sh /go/src/app/.
COPY create_database.sh /go/src/app/.
COPY migrations /go/src/app/migrations


RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.1.0

RUN apt-get update 
RUN apt-get upgrade -y
RUN apt install -y default-mysql-client

ENTRYPOINT ["./wait_for_db.sh", "./create_database.sh" ]
