package model


type TaskTouch struct{
    Id                  int32    `json:"id"`		     
    UserId              int32    `json:"userId"`	    
    TaskId              int32    `json:"taskId"`	      
    TouchTimeStamp      int64    `json:"touchTimeStamp"`     
    LocationTimeStamp   int64    `json:"locationTimeStamp"`  
    Longitude           double   `json:"longitude"`	    
    Latitude            double   `json:"latitude"`	    
    Accuracy            float32  `json:"accuracy"`	    
    NetworkType         string   `json:"networkType"`	    
    TouchType           string   `json:"touchType"`         
}
