package scheduler

import (
	"fmt"
	cron "github.com/robfig/cron/v3"
	"go-schedule/models"
	"go-schedule/services/customers"
)

func Run() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {

		fmt.Println("Run Every 1 minute on gin")
		customer := models.Customer{
			First_name: "Ragil",
			Last_name:  "Wirya",
			Phone:      "123-456-7890",
			Email:      "john@example.com",
			Street:     "123 Main St",
			City:       "Anytown",
			State:      "CA",
			Zip_code:   "12345",
		}
		customers.Create_Customer(customer)
	})
	c.Start()
}
