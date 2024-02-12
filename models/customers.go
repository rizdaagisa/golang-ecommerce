package models

type Customer struct {
	Customer_id int    `json:"customer_id"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip_code    string `json:"zip_code"`
}
