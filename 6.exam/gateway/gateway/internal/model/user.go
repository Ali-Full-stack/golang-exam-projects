package model

type UserCode struct{
	Email string `json:"email"`
	Code int64  `json:"code"` 
}

type NewPass struct {
	Email string `json:"email"`
	Code int64  `json:"code"`
	Password string `json:"password"`
}