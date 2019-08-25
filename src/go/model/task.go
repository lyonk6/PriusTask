package model

type Task struct {
	id                   int32
	userId               int32
	memo                 string
	repeatIntervalInDays int64
	taskLength           int64
	dueDate              int64
	creationDate         int64
	creationLongitude    int64
	creationLatitude     int64
}
