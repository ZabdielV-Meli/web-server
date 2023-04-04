package store

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

type jsonData struct {
	data []domain.Producto
}

// Inicializa los productos a un slice
func (repository *jsonData) InicializarBase() {
	//leer el archivo JSON
	arrayBites, err := os.ReadFile("../../products.json")
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

// All
func (repository *jsonData) GetAll() []domain.Producto {
	//leer el archivo JSON
	return repository.data
}

// Buscar por id
func (repository *jsonData) BuscarPorId(id int) (product domain.Producto, err error) {
	//leer el archivo JSON
	for _, valor := range repository.data {
		if valor.ID == id {
			return valor, nil
		}
	}

	return product, ErrprodcutNotFound
}

func (repository *jsonData) Update(product *domain.Producto) (pr domain.Producto, err error) {
	for i, valor := range repository.data {
		if valor.ID == product.ID {
			repository.data[i] = *product
			return *product, nil
		}
	}
	return pr, ErrprodcutNotFound
}

func (repository *jsonData) Delete(id int) (err error) {

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
