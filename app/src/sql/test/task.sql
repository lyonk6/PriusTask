-- // Task is an Object for holding a task.
-- type Task struct {
--     ID                   int32  `json:'id'`
--     UserID               int32  `json:'userId'`
--     Memo                 string `json:'memo'`
--     RepeatIntervalInDays int64  `json:'repeatIntervalInDays'`
--     TaskLength           int64  `json:'taskLength'`
--     DueDate              int64  `json:'dueDate'`
--     CreationDate         int64  `json:'creationDate'`
--     CreationLongitude    int64  `json:'creationLongitude'`
--     CreationLatitude     int64  `json:'creationLatitude'`
-- }
--
DROP TABLE task;
CREATE TABLE task(
 ID     BIGSERIAL PRIMARY KEY,
 UserID               BIGSERIAL,
 Memo                 VARCHAR(1000),
 RepeatIntervalInDays BIGINT,
 TaskLength           BIGINT,
 DueDate              BIGINT,
 CreationDate         BIGINT,
 CreationLongitude    BIGINT,
 CreationLatitude     BIGINT
);


INSERT INTO task(CreationDate, CreationLatitude, CreationLongitude,      DueDate,  Memo                   , RepeatIntervalInDays, TaskLength, UserId)
VALUES
(156396319123,                0,                 0,1573367057100,  'Buy Dog Food'         ,                   60,     900000,      0),
(156396319214,                0,                 0,1573367057010,  'Wash Cat'             ,                    7,     900000,      0),
(156396919214,                0,                 0,1573367057010,  'Mow Lawn'             ,                    7,     900000,      0),
(156396919214,                0,                 0,1573367057010,  'Walk Dod'             ,                    7,     900000,      0),
(156396319305,                0,                 0,1573367057000,  'File for an extension',                  365,     900000,      0),
(156396319496,                0,                 0,1563963191000,  'Check the Time'       ,                    0,     900000,      0),
(156396319587,                0,                 0, 517678260000,  'Start Life'           ,                    0,     900000,      0),
(156396319678,                0,                 0, 946684800000,  'Survive y2k'          ,                    0,     900000,      0),
(000000000001,                0,                 0,1573367057100,  'Pour (Drank)'         ,                   60,     900000,      0),
(000000000010,                0,                 0,1573367057100,  'Head shot (Drank)'    ,                   60,     900000,      0),
(000000000100,                0,                 0,1573367057100,  'Sit down (Drank)'     ,                   60,     900000,      0),
(000000001000,                0,                 0,1573367057100,  'Stand up (Drank)'     ,                   60,     900000,      0),
(000000010000,                0,                 0,1573367057100,  'Passout (Drank)'      ,                   60,     900000,      0),
(000000100000,                0,                 0,1573367057100,  'Wake up (Drank)'      ,                   60,     900000,      0),
(000001000000,                0,                 0,1573367057100,  'Faded (Drank)'        ,                   60,     900000,      0),
(000010000000,                0,                 0,1573367057100,  'Faded, faded'         ,                   60,     900000,      0),
(000100000000,                0,                 0,1573367057100,  'Get fucked up'        ,                   60,     900000,      0),
(001000000000,                0,                 0,1573367057100,  'Quit lousey job'      ,                   60,     900000,      0),
(010000000000,                0,                 0,1573367057100,  'Call recruiter'       ,                   60,     900000,      0),
(100000000000,                0,                 0,1573367057100,  'Clean counter'        ,                   60,     900000,      0);


INSERT INTO task(CreationDate, CreationLatitude, CreationLongitude,      DueDate,  Memo                   , RepeatIntervalInDays, TaskLength, UserId)
VALUES
(156396319678,                0,                 0, 946684800000,  'Return an id'         ,                    0,     900000,      0) RETURNING id;
