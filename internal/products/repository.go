package products

import "github.com/zabdielv/gin-exercises/internal/domain"

type Repository interface {
	InicializarBase()
	Save(product *domain.Producto) (err error)
	GetAll() ([]domain.Producto, error)
	BuscarPorId(id int) (product domain.Producto, err error)
}
