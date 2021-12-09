#! /bin/ash

mysql -h"db" -u "root" -p"password" -e "CREATE DATABASE world DEFAULT CHARACTER SET utf8mb4;"
/bin/migrate -database ${MYSQL_URL} -path db/migrations up

