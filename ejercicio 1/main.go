package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
}

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})

	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.POST("/saludo", func(c *gin.Context) {
		var newPerson Persona

		jsonData, _ := ioutil.ReadAll(c.Request.Body)

		if err := json.Unmarshal(jsonData, &newPerson); err != nil {
			log.Fatal(err)
			return
		}

		/* 		if err := c.BindJSON(&newPerson); err != nil {
			log.Fatal(err)
			return
		} */

		fmt.Println(newPerson)
		c.String(200, "Hola "+newPerson.Nombre+" "+newPerson.Apellido)
		fmt.Printf("Hola + %s + %s \n", newPerson.Nombre, newPerson.Apellido)

	})

	/* 	router.POST("/hello-name", func(ctx *gin.Context) {
		var personParams PersonParams
		err := ctx.BindJSON(&personParams)
		if err != nil {
			log.Fatal(err)
		}
		response := fmt.Sprintf("Hola %s %s", personParams.FirstName, personParams.LastName)
		ctx.String(http.StatusCreated, response)
	}) */
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
