package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID           int     `json:"id" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Quantity     int     `json:"quantity" binding:"required"`
	Code_value   int     `json:"code_value" binding:"required"`
	Is_published bool    `json:"is_published" binding:"required"`
	Expiration   string  `json:"expiration" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
}

func buscarPorId(id int, sliceProdcuctos *[]Producto) (Producto, error) {

	for _, valor := range *sliceProdcuctos {
		if valor.ID == id {
			return valor, nil
		}
	}

	return Producto{}, errors.New("ID no encontrado")
}

func filtrarSlice(costo float64, sliceProdcuctos *[]Producto) []Producto {

	sliceQuery := []Producto{}

	for _, valor := range *sliceProdcuctos {
		if valor.Price > costo {
			sliceQuery = append(sliceQuery, valor)
		}
	}

	return sliceQuery
}

func main() {

	//Slice de productos
	var sliceProdcuctos []Producto

	//leer el archivo JSON
	arrayBites, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println("el archivo indicado no fue encontrado")
		return
	} else {
		if err2 := json.Unmarshal(arrayBites, &sliceProdcuctos); err != nil {
			log.Fatal(err2)
			return
		}
	}

	/*
		 	fmt.Println(len(sliceProdcuctos))
			fmt.Println(sliceProdcuctos[0].Price)
	*/

	// Crea un router con gin
	router := gin.Default()

	//Ruta /ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Ruta /products
	router.GET("/products", func(c *gin.Context) {

		c.JSON(200, sliceProdcuctos)
	})

	//Ruta /products/:id
	router.GET("/products/:id", func(c *gin.Context) {

		//castear de string a int
		id, _ := strconv.Atoi(c.Param("id"))

		//Buscar por id
		producto, err := buscarPorId(id, &sliceProdcuctos)
		if err != nil {
			c.JSON(404, gin.H{
				"message": "Id no existe",
			})
		} else {
			c.JSON(200, producto)
		}

	})

	//Ruta /products/search
	router.GET("/products/", func(c *gin.Context) {

		//castear de string a int
		priceGt, _ := strconv.ParseFloat(c.Query("priceGt"), 64)

		sliceQuery := filtrarSlice(priceGt, &sliceProdcuctos)

		c.JSON(200, sliceQuery)

	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
	fmt.Println("Servidor corriendo")
}
