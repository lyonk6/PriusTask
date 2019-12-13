package model

// User is an account that owns Tasks.
type User struct {
	ID       int32  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
