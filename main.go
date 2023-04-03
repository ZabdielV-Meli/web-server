package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   int     `json:"code_value" `
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration" `
	Price        float64 `json:"price"`
}

func buscarPorId(id int, sliceProdcuctos *[]Producto) (Producto, error) {

	for _, valor := range *sliceProdcuctos {
		if valor.ID == id {
			return valor, nil
		}
	}

	return Producto{}, errors.New("ID no encontrado")
}

func filtrarSlice(c *gin.Context) {

	//castear de string a int
	priceGt, _ := strconv.ParseFloat(c.Query("priceGt"), 64)

	sliceQuery := []Producto{}

	for _, valor := range sliceProdcuctos {
		if valor.Price > priceGt {
			sliceQuery = append(sliceQuery, valor)
		}
	}

	c.JSON(200, sliceQuery)
}

func buscarProducto(c *gin.Context) {

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

}

func guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		type request struct {
			Name         string  `json:"name"`
			Quantity     int     `json:"quantity"`
			Code_value   int     `json:"code_value" `
			Is_published bool    `json:"is_published"`
			Expiration   string  `json:"expiration" `
			Price        float64 `json:"price"`
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": "parametros no validos",
			})
			fmt.Println(err)
			return
		}
		//parametros vacios
		if req.Name == "" || req.Price == 0.0 || req.Code_value == 0 || req.Quantity == 0 || req.Price == 0 || req.Expiration == "" {
			ctx.JSON(404, gin.H{
				"error": "parametro/s vacios",
			})
			return
		}
		//Code value unico
		for _, valor := range sliceProdcuctos {
			if req.Code_value == valor.Code_value {
				ctx.JSON(404, gin.H{
					"error": "code_value no es unico",
				})
				return
			}
		}

		//validar fecha
		re := regexp.MustCompile("^[0-3]?[0-9][/][0-3]?[0-9][/]([0-9]{2})?[0-9]{2}$")
		//_, err := time.Parse("01/02/2006", req.Expiration)
		if !re.MatchString(req.Expiration) {
			ctx.JSON(404, gin.H{
				"error": "fecha en formato incorrecto",
			})
			return
		}

		newProduct := Producto{
			ID:           len(sliceProdcuctos) + 1,
			Name:         req.Name,
			Quantity:     req.Quantity,
			Code_value:   req.Code_value,
			Is_published: req.Is_published,
			Expiration:   req.Expiration,
			Price:        req.Price,
		}

		//req.ID = len(sliceProdcuctos) + 1
		sliceProdcuctos = append(sliceProdcuctos, newProduct)
		ctx.JSON(200, gin.H{
			"ok":       "producto añadido correctamente",
			"producto": req,
		})
	}
}

// Slice de productos
var sliceProdcuctos []Producto

// Inicializa los productos a un slice
func getData(sliceProducts *[]Producto) {
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

}
func main() {

	// Crea un router con gin
	router := gin.Default()

	//Get data
	getData(&sliceProdcuctos)

	//Ruta /ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Ruta /products
	router.GET("/products", func(c *gin.Context) {

		c.JSON(200, sliceProdcuctos)
	})

	//Ruta /products/:id
	router.GET("/products/:id", buscarProducto)

	pr := router.Group("/products")

	//Ruta /products/search
	pr.GET("/", filtrarSlice)
	//Guardar producto
	pr.POST("/", guardar())

	//Ruta /products/search
	//router.GET("/products", filtrarSlice)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
	fmt.Println("Servidor corriendo")
}
