package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	logger "github.com/sirupsen/logrus"
	"go-schedule/controllers"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	logger.Println("Server is started ...")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/customers", controllers.Create_customer)
	router.GET("/customers", controllers.Get_all_customer)

	if router.Run("localhost:8080") != nil {
		return
	}
}

func Run() {
	c := cron.New(cron.WithChain( // Init cron with recovery panic
		cron.Recover(cron.DefaultLogger),
	))
	c.AddFunc("* * * * * *", func() {
		fmt.Println("Run Every 1 minute on gin")
	})
	c.Start()
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}
