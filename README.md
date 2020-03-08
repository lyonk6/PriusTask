# PriusTask
---
Dependencies:

  1. Golang: version 1.11

  2. Postgres for go package: https://github.com/lib/pq

  3. PostgreSQL 11.5


# Parameter File
The parameter file specifies configuration setting such as the host port, production database and test database. Use "port=", "db=" and "testdb=" to specify the port, database and test database.

```
port=
db=
testdb=
```

See also example_paramete_file

Buiding SQL files: 
```
psql --dbname=prius_task_test -f task.sql 
```



Create a self signed certificate. 
```
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
```
TODO: Make docker do this? 