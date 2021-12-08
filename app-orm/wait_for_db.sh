#! /bin/ash

set -e
cmd="$1"

while :
do
  if mysql -h"engineer-exam-mysql-host" -u "root" -p"password" -e "quit"; then
    >&2 echo "Mysql is up"
    break
  else
    >&2 echo "Mysql is unavailable - sleeping"
    sleep 1
  fi
done

exec $cmd
