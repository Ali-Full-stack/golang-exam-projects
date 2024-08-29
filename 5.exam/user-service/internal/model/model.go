package model

type UserWithID struct{
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

 type UserLogin struct{
	Id string `json:"id"`
	Password string `json:"password"`
 }

