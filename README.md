# PriusTask
---
Dependencies:

  1. Golang: version 1.11

  2. Postgres for go package: https://github.com/lib/pq

  3. PostgreSQL 11.5



Presently to test one has to have a local test psql database and a test file `test_params` which lists the test database: 
```
testdb=postgres://<user>:@localhost/prius_task_test?sslmode=disable
```

Buiding SQL files: 
```
psql --dbname=prius_task_test -f task.sql 
```
