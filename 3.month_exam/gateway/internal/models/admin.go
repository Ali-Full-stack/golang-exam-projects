package models

type AdminInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string 	`json:"role"`
}

type AdminLogin struct{
	Id string `json:"id"`
	Role string `json:"role"`
	HashPassword string `json:"hashpassword"`
}

type AdminID struct{
	Id string `json:"id"`
}

type AdminResponse struct{
	Status string `json:"status"`
}