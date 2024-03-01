package usecase

import "github.com/Patrikesm/kafka-with-go/internal/entity"

type ListProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductsRepository
}

func NewListProductUseCase(productRepository entity.ProductsRepository) *ListProductUseCase {
	return &ListProductUseCase{ProductRepository: productRepository}
}

func (u *ListProductUseCase) Execute() ([]*ListProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutput []*ListProductOutputDto

	// aqui o next não funciona pois ja estamos lidando com o objeto e não com as linhas
	for _, product := range products {
		productsOutput = append(productsOutput, &ListProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutput, nil
}
