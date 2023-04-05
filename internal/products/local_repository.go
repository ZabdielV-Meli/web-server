package products

import (
	"errors"

	"github.com/zabdielv/gin-exercises/internal/domain"
	"github.com/zabdielv/gin-exercises/pkg/store"
)

var (
	ErrprodcutNotFound = errors.New("product not found")
)

type Local_repository struct {
	//data []domain.Producto
	BD store.JsonData
}

func (data *Local_repository) GetAll() (result []domain.Producto, err error) {
	result = data.BD.GetAll()
	return
}

func (data *Local_repository) Save(product *domain.Producto) (err error) {

	// save
	err = data.BD.Save(product)

	return
}

func (data *Local_repository) BuscarPorId(id int) (product domain.Producto, err error) {

	product, err = data.BD.BuscarPorId(id)

	return

}

func (data *Local_repository) Update(product *domain.Producto) (pr domain.Producto, err error) {

	pr, ErrprodcutNotFound = data.BD.Update(product)
	return
}

func (data *Local_repository) Delete(id int) (err error) {

	err = data.BD.Delete(id)
	return
}
