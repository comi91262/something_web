#! /bin/bash

/go/bin/migrate -database "${MYSQL_URL}" -path migrations up


