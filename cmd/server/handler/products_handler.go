package handlers

import (
	"errors"
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zabdielv/gin-exercises/internal/products"
)

var (
	ErrMovieNotFound       = errors.New("movie not found")
	ErrMovieErrorArguments = errors.New("invalid or nul arguments")
)

// controller
type ProductHandler struct {
	Service products.Service
}

func (ph ProductHandler) Save() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//obtener la peticion y validarla
		var req CreateProductRequest

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
				"error": ErrMovieErrorArguments.Error(),
			})
			fmt.Println(ErrMovieErrorArguments.Error())
			return
		}

		productToCreate := req.ToDomain()
		//Crear producto
		err := ph.Service.Save(&productToCreate)

		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"ok":       "producto añadido correctamente",
			"producto": productToCreate,
		})
		//Si se quiere ocultar datos personales se crea un struct con respuesta personalizada
		/* 		ctx.JSON(200, gin.H{
			"ok":       "producto añadido correctamente",
			"producto": CreateProductResponse{
				ID: ...,
			},
		}) */
	}
}

// Obtener productos
func (ph ProductHandler) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		sliceProdcuctos, err := ph.Service.GetAll()
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, sliceProdcuctos)

	}

}

// Filtrar por precio
func (ph ProductHandler) FiltrarPorPrecio() gin.HandlerFunc {

	return func(c *gin.Context) {
		//castear de string a int
		priceGt, _ := strconv.ParseFloat(c.Query("priceGt"), 64)

		sliceProdcuctos, err := ph.Service.FiltrarPorPrecio(priceGt)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, sliceProdcuctos)

	}

}

// Buscar por ID
func (ph ProductHandler) BuscarPorId() gin.HandlerFunc {

	return func(c *gin.Context) {

		//castear de string a int
		id, _ := strconv.Atoi(c.Param("id"))

		//Buscar por id
		producto, err := ph.Service.BuscarPorId(id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, producto)
		}
	}

}
