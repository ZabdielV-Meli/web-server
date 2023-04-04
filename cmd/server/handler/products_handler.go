package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

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

func (ph ProductHandler) ValidarToken(token string) (err error) {
	tokenEnv := os.Getenv("TOKEN")
	fmt.Println("tokens: ", tokenEnv, token)
	if token != tokenEnv {
		err = errors.New("Token invalido")
	}

	return
}

func (ph ProductHandler) Save() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		header := ctx.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

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
		err = ph.Service.Save(&productToCreate)

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

func (ph ProductHandler) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// request

		header := ctx.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		//castear de string a int
		id, _ := strconv.Atoi(ctx.Param("id"))

		//Ver si existe ID
		/* 		producto, err := ph.Service.BuscarPorId(id)
		   		if err != nil {
		   			ctx.JSON(404, gin.H{
		   				"message": err.Error(),
		   			})
		   			return
		   		} */
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

		productToUpdate := req.ToDomain()
		productToUpdate.ID = id
		//Actualizar producto
		productUpdated, err := ph.Service.Update(&productToUpdate)

		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"ok":       "producto actualizado correctamente",
			"producto": productUpdated,
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

func (ph ProductHandler) Patch() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		header := ctx.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		//castear de string a int
		id, _ := strconv.Atoi(ctx.Param("id"))

		//Ver si existe ID
		producto, err := ph.Service.BuscarPorId(id)
		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		//Actualizar solo datos enviados

		if err := ctx.ShouldBindJSON(&producto); err != nil {
			ctx.JSON(404, gin.H{
				"error": "parametros no validos",
			})
			fmt.Println(err)
			return
		}

		producto.ID = id
		//Actualizar producto
		productUpdated, err := ph.Service.Update(&producto)

		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"ok":       "producto actualizado correctamente",
			"producto": productUpdated,
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

		header := c.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

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
		header := c.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
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

// Buscar por ID
func (ph ProductHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		// request
		header := c.GetHeader("TOKEN")
		err := ph.ValidarToken(header)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}
		//castear de string a int
		id, _ := strconv.Atoi(c.Param("id"))

		//Buscar por id
		// process
		err = ph.Service.Delete(id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		}

		// response
		c.Header("Location", fmt.Sprintf("/movies/%d", id))
		c.JSON(http.StatusNoContent, nil)

	}

}
