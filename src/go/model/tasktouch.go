package model


type TaskTouch struct{
    userId              int32
    taskId              int32    
    touchTimeStamp      int64   
    locationTimeStamp   int64   
    longitude           double 
    latitude            double 
    accuracy            float32
    networkType         string 
    touchType           string
}
