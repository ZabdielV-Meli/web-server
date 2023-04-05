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

type JsonData struct {
	data []domain.Producto
}

// Inicializa los productos a un slice
func (repository *JsonData) InicializarBase(ruta string) {
	//leer el archivo JSON
	arrayBites, err := os.ReadFile(ruta)
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

// Guardar
func (repository *JsonData) Save(product *domain.Producto) (err error) {
	product.ID = len(repository.data) + 1
	// save
	repository.data = append(repository.data, *product)

	return
}

// All
func (repository *JsonData) GetAll() []domain.Producto {

	//leer el archivo JSON
	//fmt.Println(repository.data)
	return repository.data
}

// Buscar por id
func (repository *JsonData) BuscarPorId(id int) (product domain.Producto, err error) {
	//leer el archivo JSON
	for _, valor := range repository.data {
		if valor.ID == id {
			return valor, nil
		}
	}

	return product, ErrprodcutNotFound
}

func (repository *JsonData) Update(product *domain.Producto) (pr domain.Producto, err error) {
	for i, valor := range repository.data {
		if valor.ID == product.ID {
			repository.data[i] = *product
			return *product, nil
		}
	}
	return pr, ErrprodcutNotFound
}

func (repository *JsonData) Delete(id int) (err error) {

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
