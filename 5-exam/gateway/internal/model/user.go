package model

type UserInfo struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
 }

 type UserLogin struct{
	Id string  `json:"id"`
	Password string `json:"password"`
 }

 type UserWithID struct{
	Id string  `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
 }