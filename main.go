package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	STATES_FILE string = "data/states.json"
)

type Application struct {
	Map *Map
}

func main() {
	app := Application{}
	app.Load(STATES_FILE)
	app.Run()
}

func (app *Application) Load(file string) {
	app.Map = &Map{}
	err := app.Map.LoadStates(file)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *Application) Run() {
	g := gin.Default()
	g.POST("/", func(c *gin.Context) {

		longitude, err := strconv.ParseFloat(c.PostForm("longitude"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Longitude Error: Argument Invalid",
			})
			return
		}

		latitude, err := strconv.ParseFloat(c.PostForm("latitude"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Latitude Error: Argument Invalid",			
			})
			return
		}

		states, err := app.Map.Contains(latitude, longitude)
		if err != nil {
			c.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		} else {
			c.JSON(http.StatusOK, states)
		}
	})

	g.Run(":8080")
}
