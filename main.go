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

		longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Errorf("Longitude Error:", err),
			})
		}

		latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Errorf("Latitude Error:", err),
			})
		}

		state, err := app.Map.Contains(latitude, longitude)
		if err != nil {
			c.Data(http.StatusNoContent, gin.MIMEHTML, nil)
		} else {
			c.JSON(http.StatusOK, []string{state.Name})
		}
	})

	g.Run(":8080")
}
