# // User is an account that owns Tasks.
CREATE TABLE account(
   ID serial PRIMARY KEY,
   Email VARCHAR (355) UNIQUE NOT NULL,
   Password VARCHAR(1048) NOT NULL
);



# // Task is an Object for holding a task.
CREATE TABLE task(
 ID        serial PRIMARY KEY,
 UserID               INTEGER,  
 Memo                 VARCHAR(1000),
 RepeatIntervalInDays BIGINT,  
 TaskLength           BIGINT,  
 DueDate              BIGINT,  
 CreationDate         BIGINT,  
 CreationLongitude    BIGINT,  
 CreationLatitude     BIGINT
);

 
# //TaskTouch is an instance of a user updating or interacting with a Task.
CREATE TABLE tasktouch(
  ID    serial PRIMARY KEY,
  UserID            INTEGER,   
  TaskID            INTEGER,   
  TouchTimeStamp    BIGINT,   
  LocationTimeStamp BIGINT,   
  Longitude         float8, 
  Latitude          float8, 
  Accuracy          float8, 
  NetworkType       varchar(100),
  TouchType         varchar(40)
);