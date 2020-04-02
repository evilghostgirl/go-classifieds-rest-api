#!/usr/bin/env bash
curl -d '{ "Name": "postgres", \
"Address": "postgres", \
"Check": { \
    "id": "mem-util", \
    "name": "Memory utilization",\
    "args": ["/usr/local/bin/check_postgres-2.25.0/check_postgres.pl", "--host=localhost", "--dbname=postgres" ,"--dbuser=hazelina", "--dbpass=secret", "--action=connection"],\
      "interval": "10s",\
      "timeout": "1s"\
    }}' -H "Content-Type: application/json; charset=utf-8" -X PUT http://consulz:8500/v1/agent/service/register;