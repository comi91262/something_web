#! /bin/bash

set -e
cmd="$1"

while :
do
  if mysql -h"db" -u "root" -p"password" -e "quit"; then
    >&2 echo "Mysql is up"
    break
  else
    >&2 echo "Mysql is unavailable - sleeping"
    sleep 1
  fi
done

exec $cmd
