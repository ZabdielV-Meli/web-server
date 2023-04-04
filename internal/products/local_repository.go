package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/zabdielv/gin-exercises/internal/domain"
)

var (
	ErrprodcutNotFound = errors.New("product not found")
)

type Local_repository struct {
	data []domain.Producto
}

// Inicializa los productos a un slice
func (repository *Local_repository) InicializarBase() {
	//leer el archivo JSON
	arrayBites, err := os.ReadFile("../products.json")
	if err != nil {
		fmt.Println("el archivo indicado no fue encontrado")
		return
	} else {
		if err2 := json.Unmarshal(arrayBites, &repository.data); err != nil {
			log.Fatal(err2)
			return
		}
	}
}
func (repository *Local_repository) GetAll() (result []domain.Producto, err error) {
	result = repository.data
	return
}

func (repository *Local_repository) Save(product *domain.Producto) (err error) {
	product.ID = len(repository.data) + 1
	// save
	repository.data = append(repository.data, *product)

	return
}

func (repository *Local_repository) BuscarPorId(id int) (product domain.Producto, err error) {

	for _, valor := range repository.data {
		if valor.ID == id {
			return valor, nil
		}
	}

	return product, ErrprodcutNotFound
}

func (repository *Local_repository) Update(product *domain.Producto) (pr domain.Producto, err error) {
	for i, valor := range repository.data {
		if valor.ID == product.ID {
			repository.data[i] = *product
			return *product, nil
		}
	}
	return pr, ErrprodcutNotFound
}

func (repository *Local_repository) Delete(id int) (err error) {

	for i, m := range repository.data {
		if m.ID == id {
			// delete movie
			repository.data = append(repository.data[:i], repository.data[i+1:]...)
			return
		}
	}
	err = ErrprodcutNotFound
	return
}
