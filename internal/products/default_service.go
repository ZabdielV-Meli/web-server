package products

import (
	"errors"
	"regexp"

	"github.com/zabdielv/gin-exercises/internal/domain"
)

var (
	ErrMovieErrorArguments = errors.New("invalid or nil arguments")
	InvalidDate            = errors.New("invalid date")
)

type DefaultService struct {
	Repository Repository
}

func (s DefaultService) Save(product *domain.Producto) (err error) {

	sliceDatos, _ := s.Repository.GetAll()
	//Code value unico
	for _, valor := range sliceDatos {
		if product.Code_value == valor.Code_value {
			return ErrMovieErrorArguments
		}
	}

	//validar fecha
	re := regexp.MustCompile("^[0-3]?[0-9][/][0-3]?[0-9][/]([0-9]{2})?[0-9]{2}$")
	//_, err := time.Parse("01/02/2006", req.Expiration)
	if !re.MatchString(product.Expiration) {
		return InvalidDate
	}
	err = s.Repository.Save(product)
	if err != nil {
		//manejar error
	}
	return
}

func (s DefaultService) BuscarPorId(id int) (product domain.Producto, err error) {

	product, err = s.Repository.BuscarPorId(id)
	if err != nil {
		//manejar error
	}
	return
}

// Obtener DB
func (s DefaultService) GetAll() (products []domain.Producto, err error) {
	products, err = s.GetAll()
	if err != nil {
		//manejar error
	}
	return
}

// Filtrar por precio
func (s DefaultService) FiltrarPorPrecio(precio float64) ([]domain.Producto, error) {

	sliceDatos, _ := s.Repository.GetAll()

	sliceQuery := []domain.Producto{}

	for _, valor := range sliceDatos {
		if valor.Price > precio {
			sliceQuery = append(sliceQuery, valor)
		}
	}

	return sliceQuery, nil
}
