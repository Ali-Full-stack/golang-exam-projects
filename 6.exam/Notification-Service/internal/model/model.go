package model

type Report struct {
	Income     float64 `json:"income" bson:"income"`
	Expenses   float64 `json:"expenses" bson:"expenses"`
	NetSavings float64 `json:"net_savings" bson:"net_savings"`
}

type Notify struct {
	UserID  string `json:"user_id" bson:"user_id"`
	Message string `json:"message" bson:"message"`
	Report  *Report `json:"report" bson:"report"`
	Seen 	bool   `json:"seen" bson:"seen"`
	Date    string `json:"date" bson:"date"`
}

type NotifyList struct {
	NotifyList []Notify `json:"notify_list" bson:"notify_list"`
}
