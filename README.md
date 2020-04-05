# PriusTask
---
Dependencies:

  1. Golang: version 1.11

  2. Postgres for go package: https://github.com/lib/pq

  3. PostgreSQL 11.5



## API
API            | Type           | Return
-------------- | -------------- | --------------
/PostTaskTouch | TaskTouch      | 200, Task[]*
/PutTask       | Task           | 200
/PostTask      | Task           | 200          

To get a list of tasks to do call `/PostTaskTouch`. If a TaskTouch object is marked as `COMPLETED`, `DISMISSED` or `START_UP` an array of 20 tasks is returned. Otherwise and an empyty task list is returned.


## Database
Presently to test one has to have a local test psql database and a test file `test_params` which lists the test database: 
```
testdb=postgres://<user>:@localhost/prius_task_test?sslmode=disable
```

Use the sql files located in app/src/sql to create mock and production databases.

Buiding SQL files: 
```
psql --dbname=prius_task_test -f task.sql 
```
