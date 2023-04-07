package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	handlers "github.com/zabdielv/gin-exercises/cmd/server/handler"
	"github.com/zabdielv/gin-exercises/internal/domain"
	"github.com/zabdielv/gin-exercises/internal/products"
	"github.com/zabdielv/gin-exercises/pkg/store"
	"github.com/zabdielv/gin-exercises/pkg/web"
)

func createServerForTestProductsHandler() *gin.Engine {

	//Variables de entorno
	err := godotenv.Load("../../.env")

	if err != nil {

		log.Fatal("Error al intentar cargar archivo .env")

	}

	// Crea un router con gin
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Iniciazilar BD
	db := store.JsonData{}
	db.InicializarBase("../../products.json")

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

	// Corremos nuestro servidor sobre el puerto 8080
	// run

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "1234")

	return req, httptest.NewRecorder()
}

// GET /products
func Test_GetProducts_OK(t *testing.T) {

	var objRes web.SuccessfulResponse
	var objProduct []domain.Producto
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	//Convertir struct SuccessfulResponse.Data a bytes
	producsResponseBytes, _ := json.Marshal(objRes.Data)
	//Convertir bytes a CreateProductsResponse
	json.Unmarshal(producsResponseBytes, &objProduct)

	//t.Log(productsResponse.Datos)
	assert.True(t, len(objProduct) > 0)
}

// GET /products/:id
func Test_GetProduct_OK(t *testing.T) {

	//ARRANGE
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()

	//ACT
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/2", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	t.Log("\ndatos del response", rr)
	//ASSERT
	assert.Equal(t, 200, rr.Code)
}

// POST /products
func Test_SaveProduct_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/products/", `{
		"name": "Chicken - Soup Base",
			"quantity": 479,
			"code_value": 45354334,
			"is_published": false,
			"expiration": "11/05/2021",
			"price": 515.93
	  }`)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code)
}

// Delete /products
func Test_DeleteProduct_OK(t *testing.T) {

	//ARRANGE
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()

	//ACT
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/products/2", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	//t.Log("\ndatos del response", rr)
	//ASSERT
	assert.Equal(t, 204, rr.Code)
}

// errores
// GET /products/sadas
func Test_GetProduct_400(t *testing.T) {

	//ARRANGE
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()

	//ACT
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/sadas", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	t.Log("\ndatos del response", rr)
	//ASSERT
	assert.Equal(t, 400, rr.Code)

}

// errores
// GET /products/404
func Test_GetProduct_404(t *testing.T) {

	//ARRANGE
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()

	//ACT
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/501", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	t.Log("\ndatos del response", rr)
	//ASSERT
	assert.Equal(t, 404, rr.Code)

}

// errores
// Delete /products
func Test_DeleteProduct_401(t *testing.T) {

	//ARRANGE
	// crear el Server y definir las Rutas
	r := createServerForTestProductsHandler()

	//ACT
	// crear Request del tipo GET y Response para obtener el resultado
	//req, rr := createRequestTest(http.MethodDelete, "/products/2", "")
	req := httptest.NewRequest(http.MethodDelete, "/products/2", bytes.NewBuffer([]byte("")))
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	//t.Log("\ndatos del response", rr)
	//ASSERT
	assert.Equal(t, 401, rr.Code)
}
