--
-- var touchTypes = []string{'SAVED', 'DELETED', 'COMPLETED', 'DISMISSED', 'START_UP', 'HEART_BEAT', 'CREATED'}
--
-- //TaskTouch is an instance of a user updating or interacting with a Task.
-- type TaskTouch struct {
--     ID                int32   `json:'id'`
--     UserID            int32   `json:'userId'`
--     TaskID            int32   `json:'taskId'`
--     TouchTimeStamp    int64   `json:'touchTimeStamp'`
--     LocationTimeStamp int64   `json:'locationTimeStamp'`
--     Longitude         float64 `json:'longitude'`
--     Latitude          float64 `json:'latitude'`
--     Accuracy          float64 `json:'accuracy'`
--     NetworkType       string  `json:'networkType'`
--     TouchType         string  `json:'touchType'`
-- }
--
DROP TABLE tasktouch;
CREATE TABLE tasktouch(
  ID  BIGSERIAL PRIMARY KEY,
  UserID            BIGSERIAL,
  TaskID            BIGSERIAL,
  TouchTimeStamp    BIGINT,
  LocationTimeStamp BIGINT,
  Longitude         float8,
  Latitude          float8,
  Accuracy          float8,
  NetworkType       varchar(100),
  TouchType         varchar(40)
);


INSERT INTO tasktouch(Accuracy,   Latitude, LocationTimeStamp,    Longitude, NetworkType, TaskId, TouchTimeStamp,   TouchType, userId)
VALUES  (20.867,   36.0008196,     1570311185403, -115.1304113,   'network',      0,  1570311194428,     'SAVED',      0),
        (20.867,   36.0008196,     1570311185403, -115.1304113,   'network',      0,  1570311208168, 'COMPLETED',      0),
        (20.867,   36.0008196,     1570311185403, -115.1304113,   'network',      1,  1570311223972,   'DELETED',      0) RETURNING ID;
