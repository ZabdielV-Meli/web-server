package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	handlers "github.com/zabdielv/gin-exercises/cmd/server/handler"
	"github.com/zabdielv/gin-exercises/internal/products"
)

func main() {

	// Crea un router con gin
	router := gin.Default()

	//Create repository
	repository := products.Local_repository{}
	repository.InicializarBase()

	//Create service
	service := products.DefaultService{
		Repository: &repository,
	}
	//Create handler
	handler := handlers.ProductHandler{
		Service: service,
	}
	//Ruta /ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Ruta /products
	router.GET("/products", handler.GetAll())

	//Ruta /products/:id
	router.GET("/products/:id", handler.BuscarPorId())

	pr := router.Group("/products")

	//Ruta /products/search
	pr.GET("/", handler.FiltrarPorPrecio())
	//Guardar producto
	pr.POST("/", handler.Save())

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
	fmt.Println("Servidor corriendo")
}
