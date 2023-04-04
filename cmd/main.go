package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	handlers "github.com/zabdielv/gin-exercises/cmd/server/handler"
	"github.com/zabdielv/gin-exercises/internal/products"
)

func main() {

	//Variables de entorno
	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error al intentar cargar archivo .env")

	}

	usuario := os.Getenv("MY_USER")

	password := os.Getenv("MY_PASS")

	println("Usuario sacado de variables de Entorno: ", usuario)

	println("Password sacado de variables de Entorno: ", password)

	// Crea un router con gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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
	router.PUT("/products/:id", handler.Update())
	router.PATCH("/products/:id", handler.Patch())
	router.DELETE("/products/:id", handler.Delete())

	pr := router.Group("/products")

	//Ruta /products/search
	pr.GET("/", handler.FiltrarPorPrecio())
	//Guardar producto
	pr.POST("/", handler.Save())

	// Corremos nuestro servidor sobre el puerto 8080
	// run

	router.Run()
	fmt.Println("Servidor corriendo")
}
