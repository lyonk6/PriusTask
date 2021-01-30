# PriusTask
---
Dependencies:

  1. Golang: version 1.11

  2. Postgres for go package: https://github.com/lib/pq

  3. PostgreSQL 11.5



## API
Method         | Type           | Return
-------------- | -------------- | --------------
/PostTaskTouch | TaskTouch      | 200, Task[]*
/PutTask       | Task           | 200
/PostTask      | Task           | 200          

To get a list of tasks, use `/PostTaskTouch`. If a TaskTouch object is marked as `COMPLETED`, `DISMISSED` or `START_UP` an array of 20 tasks is returned. Otherwise and an empty task list is returned.


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

## Using the API:

###Start up:
When the applications starts make a call /PostTaskTouch to retrieve the current task list:

```
Address: https://127.0.0.0:1234/PostTaskTouch
Body: {
    "accuracy": 0.0,
    "id": 0,
    "latitude": 0.0,
    "locationTimeStamp": 0,
    "longitude": 0.0,
    "taskId": -1,
    "touchTimeStamp": 1600729082413,
    "touchType": "START_UP",
    "userId": 0
}
```

The field "touchType" describes why we are reaching out to the API. In the case above, we are simply starting the application so we pass the value "START_UP". This tells the API to pull a task list based on the current time and location values in the taskTouch object.

Any time a task is created, updated, deleted, completed, or dismissed an appropriate TaskTouch object should be sent to the API using the `/PostTaskTouch` method.

Valid task touch types include "UPDATED", "DELETED", "COMPLETED", "DISMISSED", "START_UP" and "CREATED".

###Getting tasks:
The only way to retrieve tasks is by calling `/PostTaskTouch`. Tasks will only be returned however if the "touchType" field is marked as `COMPLETED`, `DISMISSED` or `START_UP`. It is not possible to retrieve a particular task, only a list of tasks the AI chooses.

###Creating tasks:
To create a new Task use the `/PostTask` method. When doing this, set the taskId=-1. Upon a successful call, the task is returned with it's new taskId.

###Updating a Task:
To update a task It is only necessary to use `/PutTask` if the a change is made to the task "memo", "dueDate", "taskLength", or "repeatIntervalInDays" fields.

```
Address: https://127.0.0.0:1234/PostTask
Body: {
	"id": -1,
	"userId": -1,
	"memo": "Bicep Curls \u0026 Pullups",
	"repeatIntervalInDays": 7,
	"taskLength": 900000,
	"dueDate": 1612022560000,
	"creationDate": 1611922529000,
	"creationLongitude": 0,
	"creationLatitude": 0,
	"lastTouchType": "CREATED"
}
```
B