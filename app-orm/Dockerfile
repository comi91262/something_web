FROM alpine:3.14

COPY wait_for_db.sh /root/wait_for_db.sh

RUN apk update
RUN apk add mysql-client

CMD ["/root/wait_for_db.sh"]
