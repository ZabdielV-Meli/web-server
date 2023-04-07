package handlers

import (
	"errors"
	"fmt"
	"strings"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zabdielv/gin-exercises/internal/products"
	"github.com/zabdielv/gin-exercises/pkg/web"
)

var (
	ErrMovieNotFound       = errors.New("movie not found")
	ErrMovieErrorArguments = errors.New("invalid or nul arguments")
)

// controller
type ProductHandler struct {
	Service products.Service
}

// funcion que trae todos los products
// comentario propio que no afecta a la documentacion
// @Summary Save a product
// @Tags Products
// @Description Save a product through body
// @Produce json
// @Accept			json
// @Param			product	body		handlers.CreateProductRequest	true	"Some product"
// @Param token header string true "TOKEN"
// @Sucess 200 {object}	web.SuccessfulResponse
// @Failure 404
// @Router /products/ [POST]
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

		/* 		gin.H{
			"ok":       "producto añadido correctamente",
			"producto": productToCreate,
		} */
		web.Success(ctx, 201, CreateProductResponse{
			ID:    productToCreate.ID,
			Name:  productToCreate.Name,
			Price: productToCreate.Price,
		})
		//web.Success(ctx, 201, productToCreate)
		//Si se quiere ocultar datos personales se crea un struct con respuesta personalizada
		/* 		ctx.JSON(200, gin.H{
			"ok":       "producto añadido correctamente",
			"producto": CreateProductResponse{
				ID: ...,
			},
		}) */
	}
}

// funcion que trae todos los products
// comentario propio que no afecta a la documentacion
// @Summary Update a product
// @Tags Products
// @Description Update a product through body
// @Produce json
// @Accept			json
// @Param id path string true "Please give ID"
// @Param product body handlers.CreateProductRequest	true	"Some product"
// @Param token header string true "TOKEN"
// @Sucess 200 {object}	web.SuccessfulResponse
// @Failure 404
// @Router /products/{id} [PUT]
func (ph ProductHandler) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// request

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

		/* 		gin.H{
			"ok":       "producto actualizado correctamente",
			"producto": productUpdated,
		} */

		web.Success(ctx, 200, CreateProductResponse{
			ID:    productUpdated.ID,
			Name:  productUpdated.Name,
			Price: productToUpdate.Price,
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

// @Summary Patch a product
// @Tags Products
// @Description Update some atributes of Model
// @Produce json
// @Accept			json
// @Param id path string true "Please give ID"
// @Param product body handlers.CreateProductRequest	true	"Some product"
// @Param token header string true "TOKEN"
// @Sucess 200 {object}	web.SuccessfulResponse
// @Failure 404
// @Router /products/{id} [PUT]
func (ph ProductHandler) Patch() gin.HandlerFunc {

	return func(ctx *gin.Context) {

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

		//redundante ?
		producto.ID = id
		//Actualizar producto
		productUpdated, err := ph.Service.Update(&producto)

		if err != nil {
			ctx.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		/* 		ctx.JSON(200, gin.H{
			"ok":       "producto actualizado correctamente",
			"producto": productUpdated,
		}) */

		web.Success(ctx, 200, CreateProductResponse{
			ID:    productUpdated.ID,
			Name:  productUpdated.Name,
			Price: productUpdated.Price,
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

// funcion que trae todos los products
// comentario propio que no afecta a la documentacion
// @Summary List Products
// @Tags Products
// @Description Gets all Products without filter
// @Produce json
// @Param token header string true "TOKEN"
// @Sucess 200 {object}	web.SuccessfulResponse
// @Failure 400
// @Router /products/ [GET]
// Obtener productos
func (ph ProductHandler) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		//Si existe un query
		valorQuery := c.Query("priceGt")

		if valorQuery != "" {
			//castear de string a int
			priceGt, _ := strconv.ParseFloat(valorQuery, 64)
			sliceProdcuctos, err := ph.Service.FiltrarPorPrecio(priceGt)
			if err != nil {
				c.JSON(200, gin.H{
					"error": err.Error(),
				})
				return
			}

			//c.JSON(200, sliceProdcuctos)

			web.Success(c, 200, sliceProdcuctos)
		}
		sliceProdcuctos, err := ph.Service.GetAll()
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		/* 	c.JSON(200, sliceProdcuctos) */

		web.Success(c, 200, sliceProdcuctos)
	}

}

// lista
func (ph ProductHandler) ListaProductos() gin.HandlerFunc {

	return func(c *gin.Context) {

		//castear de string a int
		//Si existe un query
		valorQuery := c.Query("list")

		//manejar error
		if valorQuery == "" || len(valorQuery) < 3 {

			c.JSON(201, gin.H{
				"error": "mal query",
			})
			return

		}

		valorQuery = valorQuery[1 : len(valorQuery)-1]
		arrayQuery := strings.Split(valorQuery, ",")

		//Convertir cada character a un int y pasarlo a un slice
		temporalSlice := []int{}
		for _, valor := range arrayQuery {
			id, err := strconv.Atoi(valor)
			if err != nil {
				c.JSON(201, gin.H{
					"error": "mal query",
				})
				return
			}
			temporalSlice = append(temporalSlice, id)
		}

		sliceProdcuctos, _, err := ph.Service.ListaProductos(temporalSlice)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		web.Success(c, 200, *sliceProdcuctos)
	}

}

// @Summary Find by ID a product
// @Tags Products
// @Description Find by ID a product
// @Produce json
// @Param id path string true "Please give ID"
// @Param token header string true "TOKEN"
// @Sucess 200 {object}	web.SuccessfulResponse
// @Failure 404
// @Router /products/{id} [GET]
// Buscar por ID
func (ph ProductHandler) BuscarPorId() gin.HandlerFunc {

	return func(c *gin.Context) {

		//castear de string a int
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": errors.New("Bad request"),
			})
			return
		}
		//Buscar por id
		producto, err := ph.Service.BuscarPorId(id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		/* 	c.JSON(200, producto) */

		web.Success(c, 200, CreateProductResponse{
			ID:    producto.ID,
			Name:  producto.Name,
			Price: producto.Price,
		})
	}

}

// @Summary Delete by ID a product
// @Tags Products
// @Description Delete by ID a product
// @Param id path string true "Please give ID"
// @Param token header string true "TOKEN"
// @Sucess 204 {object}
// @Failure 404
// @Router /products/{id} [delete]
// Buscar por ID
// Buscar por ID
func (ph ProductHandler) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		// request

		//castear de string a int
		id, _ := strconv.Atoi(c.Param("id"))

		//Buscar por id
		// process
		err := ph.Service.Delete(id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
		}

		// response
		c.Header("Location", fmt.Sprintf("/movies/%d", id))

		web.Success(c, 204, nil)

	}

}
