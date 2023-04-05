package products

import "github.com/zabdielv/gin-exercises/internal/domain"

type Service interface {
	Save(product *domain.Producto) (err error)
	GetAll() ([]domain.Producto, error)
	FiltrarPorPrecio(precio float64) ([]domain.Producto, error)
	BuscarPorId(id int) (product domain.Producto, err error)
	Update(product *domain.Producto) (pr domain.Producto, err error)
	Delete(id int) error
	ListaProductos(sliceIds []int) (sliceProductos *[]domain.Producto, totalImpuesto float64, err error)
}
