package model


type User struct{
    Id         int32    `json:"id"`		     
    Email      string   `json:"email"`	    
    Password   string   `json:"password"`	      
}
