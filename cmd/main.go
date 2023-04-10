package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "github.com/zabdielv/gin-exercises/docs"

	// swagger embed files

	"github.com/joho/godotenv"
	handlers "github.com/zabdielv/gin-exercises/cmd/server/handler"
	"github.com/zabdielv/gin-exercises/internal/products"
	"github.com/zabdielv/gin-exercises/pkg/store"
)

// @title           Products API
// @version         1.0
// @description     APi example
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /products/

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	//Variables de entorno
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("GG")
		log.Fatal("Error al intentar cargar archivo .env")

	}

	// Crea un router con gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Iniciazilar BD
	db := store.JsonData{}
	db.InicializarBase("../products.json")

	//Create repository
	repository := products.Local_repository{
		BD: db,
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

	router.GET("/products/consumer_price/", handler.ListaProductos())
	//Ruta /products/:id
	router.GET("/products/:id", handler.BuscarPorId())
	router.PUT("/products/:id", handler.Update())
	router.PATCH("/products/:id", handler.Patch())
	router.DELETE("/products/:id", handler.Delete())

	pr := router.Group("/products/")

	pr.GET("/", handler.GetAll())
	//Guardar producto
	pr.POST("/", handler.Save())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Corremos nuestro servidor sobre el puerto 8080
	// run

	router.Run()
	fmt.Println("Servidor corriendo")
}
