#! /bin/bash

mysql -h"db" -u "root" -p"password" -e "CREATE DATABASE world DEFAULT CHARACTER SET utf8mb4;"
/go/bin/migrate -database "${MYSQL_URL}" -path migrations up


