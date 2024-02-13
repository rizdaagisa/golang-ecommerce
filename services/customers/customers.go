package customers

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-schedule/DB"
	"go-schedule/models"
	"go-schedule/utils"
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
		sql.Named("ZipCode", customer.Zip_code),
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func Export_DB() error {
	query := `SELECT * FROM sales.customers`

	connection := DB.Connect()
	if connection == nil {
		panic("error")
	}

	rows, err := connection.Query(query)
	if err != nil {
		return err // Return both nil slice and the error
	}
	defer rows.Close()

	var data_customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		var phone sql.NullString // Use sql.NullString to handle NULL values

		err = rows.Scan(&customer.Customer_id, &customer.First_name, &customer.Last_name, &phone, &customer.Email, &customer.Street, &customer.City, &customer.State, &customer.Zip_code)
		if err != nil {
			return err // Return both nil slice and the error
		}

		if phone.Valid {
			customer.Phone = phone.String // Use phone.String if the phone value is not NULL
		}

		data_customers = append(data_customers, customer)
	}

	err = utils.Export_DB(data_customers, "customers", "customers", "", "xlsx")
	if err != nil {
		return err // Return both nil slice and the error
	}

	return nil
}

func Import_DB() error {

	xlsx, err := excelize.OpenFile("upload.xlsx")

	if err != nil {
		fmt.Println("Error opening Excel file:", err)
		return err
	}

	rows, _ := xlsx.GetRows("Data export_customer")

	for i, row := range rows {
		if i == 0 { // Skip header row
			continue
		}

		for _, cell := range row {
			if len(cell) > 15 {
				// Handle the oversized data (e.g., truncate, report an error)
				fmt.Printf("Error: Data exceeds maximum length for column: %s\n", cell)
				// Optionally, you can return an error here if you want to stop processing the current row
			}
		}

		customer := models.Customer{
			//Customer_id: parseToInt(row[0]),
			First_name: row[1],
			Last_name:  row[2],
			Phone:      row[3],
			Email:      row[4],
			Street:     row[5],
			City:       row[6],
			State:      row[7],
			Zip_code:   row[8],
		}

		fmt.Println(customer)

		err := Create_Customer(customer)
		if err != nil {
			fmt.Println("Error inserting data into database at :", i, err)
			return err
		}
	}

	fmt.Println("Data imported successfully!")

	return nil
}

func Export_DB2() error {
	query := `SELECT * FROM sales.customers`

	rows, err := connection.Query(query)
	if err != nil {
		return err // Return both nil slice and the error
	}
	defer rows.Close()

	var data_customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		var phone sql.NullString // Use sql.NullString to handle NULL values

		err := rows.Scan(&customer.Customer_id, &customer.First_name, &customer.Last_name, &phone, &customer.Email, &customer.Street, &customer.City, &customer.State, &customer.Zip_code)
		if err != nil {
			return err // Return both nil slice and the error
		}

		if phone.Valid {
			customer.Phone = phone.String // Use phone.String if the phone value is not NULL
		}

		data_customers = append(data_customers, customer)
	}

	if err := rows.Err(); err != nil {
		return err // Return both nil slice and the error
	}

	file := excelize.NewFile()

	file.SetSheetName(file.GetSheetName(0), "Data export_customer")
	file.SetCellValue("Data export_customer", "A1", "ID customer")
	file.SetCellValue("Data export_customer", "B1", "First Name")
	file.SetCellValue("Data export_customer", "C1", "Last Name")
	file.SetCellValue("Data export_customer", "D1", "Phone")
	file.SetCellValue("Data export_customer", "E1", "Email")
	file.SetCellValue("Data export_customer", "F1", "Street")
	file.SetCellValue("Data export_customer", "G1", "City")
	file.SetCellValue("Data export_customer", "H1", "State")
	file.SetCellValue("Data export_customer", "I1", "Zip Code")

	for i, each := range data_customers {
		file.SetCellValue("Data export_customer", fmt.Sprintf("A%d", i+2), each.Customer_id)
		file.SetCellValue("Data export_customer", fmt.Sprintf("B%d", i+2), each.First_name)
		file.SetCellValue("Data export_customer", fmt.Sprintf("C%d", i+2), each.Last_name)
		file.SetCellValue("Data export_customer", fmt.Sprintf("D%d", i+2), each.Email)
		file.SetCellValue("Data export_customer", fmt.Sprintf("E%d", i+2), each.Phone)
		file.SetCellValue("Data export_customer", fmt.Sprintf("F%d", i+2), each.Street)
		file.SetCellValue("Data export_customer", fmt.Sprintf("G%d", i+2), each.City)
		file.SetCellValue("Data export_customer", fmt.Sprintf("H%d", i+2), each.State)
		file.SetCellValue("Data export_customer", fmt.Sprintf("I%d", i+2), each.Zip_code)
	}

	errs := file.SaveAs("customers.xlsx")
	if errs != nil {
		// Handle error
		fmt.Println("Failed to save Excel file:", err)
		return errs
	}

	return nil // Return the slice and no error
}
