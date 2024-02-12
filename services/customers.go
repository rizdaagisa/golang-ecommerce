package customers

import (
	"database/sql"
	"fmt"
	"go-schedule/DB"
	"go-schedule/models"
)

var connection *sql.DB

func init() {
	connection = DB.Connect()
	if connection == nil {
		panic("error")
	}
}

func Get_all() ([]models.Customer, error) {
	query := `SELECT * FROM sales.customers`

	rows, err := connection.Query(query)
	if err != nil {
		return nil, err // Return both nil slice and the error
	}
	defer rows.Close()

	var data_customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		var phone sql.NullString // Use sql.NullString to handle NULL values

		err := rows.Scan(&customer.Customer_id, &customer.First_name, &customer.Last_name, &phone, &customer.Email, &customer.Street, &customer.City, &customer.State, &customer.Zip_code)
		if err != nil {
			return nil, err // Return both nil slice and the error
		}

		if phone.Valid {
			customer.Phone = phone.String // Use phone.String if the phone value is not NULL
		}

		data_customers = append(data_customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err // Return both nil slice and the error
	}

	return data_customers, nil // Return the slice and no error
}

func Create_Customer(customer models.Customer) error {

	query := `
	   INSERT INTO sales.customers
	       (first_name, last_name, phone, email, street, city, state, zip_code)
	   VALUES
	       (@FirstName, @LastName, @Phone, @Email, @Street, @City, @State, @ZipCode)
	`

	_, err := connection.Exec(query,
		//sql.Named("Customer_id", nil),
		sql.Named("FirstName", customer.First_name),
		sql.Named("LastName", customer.Last_name),
		sql.Named("Phone", customer.Phone),
		sql.Named("Email", customer.Email),
		sql.Named("Street", customer.Street),
		sql.Named("City", customer.City),
		sql.Named("State", customer.State),
		sql.Named("ZipCode", customer.Last_name),
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
