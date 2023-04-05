package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	handlers "github.com/zabdielv/gin-exercises/cmd/server/handler"
	"github.com/zabdielv/gin-exercises/internal/products"
	"github.com/zabdielv/gin-exercises/pkg/store"
)

func main() {

	//Variables de entorno
	err := godotenv.Load("../.env")

	if err != nil {

		log.Fatal("Error al intentar cargar archivo .env")

	}

	// Crea un router con gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Iniciazilar BD
	data := store.JsonData{}
	data.InicializarBase()

	//Create repository
	repository := products.Local_repository{
		BD: data,
	}

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
	//router.GET("/products", handler.GetAll())

	router.GET("/products/consumer_price", handler.ListaProductos())
	//Ruta /products/:id
	//router.GET("/products/:id", handler.BuscarPorId())
	router.PUT("/products/:id", handler.Update())
	router.PATCH("/products/:id", handler.Patch())
	router.DELETE("/products/:id", handler.Delete())

	pr := router.Group("/products/")

	pr.GET("/", handler.GetAll())
	//Guardar producto
	pr.POST("/", handler.Save())

	// Corremos nuestro servidor sobre el puerto 8080
	// run

	router.Run()
	fmt.Println("Servidor corriendo")
}
