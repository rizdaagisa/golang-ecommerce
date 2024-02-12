package customers

import (
	"database/sql"
	"go-schedule/DB"
	"go-schedule/models"
)

var connection *sql.DB

func init() {
	connection = DB.GetConnection()
	if connection == nil {
		panic("error")
	}
}

func Get_all() []models.Customer {
	query := `select * from sales.customers`

	rows, err := connection.Query(query)
	if err != nil {
		panic("error")
	}
	defer rows.Close()

	var data_customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		rows.Scan(&customer.Customer_id, &customer.First_name, &customer.Last_name, &customer.Phone, &customer.Email, &customer.Street, &customer.City, &customer.State, &customer.Zip_code)
		data_customers = append(data_customers, customer)
	}

	return data_customers
}

func Create_Customer(customer models.Customer) error {
	query := "INSERT INTO sales.customers (Customer_id, First_name, Last_name, Phone, Email, Street, City, State, Zip_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := connection.Exec(query, customer.Customer_id, customer.First_name, customer.Last_name, customer.Phone, customer.Email, customer.Street, customer.City, customer.State, customer.Zip_code)

	if err != nil {
		return err
	}

	return nil
}
