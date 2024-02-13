package staffs

import (
	"database/sql"
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

func Export_DB() error {
	query := `SELECT staff_id, first_name, last_name, email, phone, active FROM sales.staffs`

	connection := DB.Connect()
	if connection == nil {
		panic("error")
	}

	rows, err := connection.Query(query)
	if err != nil {
		return err // Return both nil slice and the error
	}
	defer rows.Close()

	var data []models.Staff

	for rows.Next() {
		var staff models.Staff
		var phone sql.NullString // Use sql.NullString to handle NULL values
		err := rows.Scan(&staff.Staff_id, &staff.First_name, &staff.Last_name, &staff.Email, &staff.Phone, &staff.Active)
		if err != nil {
			return err // Return both nil slice and the error
		}

		if phone.Valid {
			staff.Phone = phone.String // Use phone.String if the phone value is not NULL
		}

		data = append(data, staff)
	}

	err = utils.Export_DB(data, "staff", "staff", "", "xlsx")

	if err != nil {
		return err
	}

	return nil
}
