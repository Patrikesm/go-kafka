package usecase

import "github.com/Patrikesm/kafka-with-go/internal/entity"

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductsRepository
}

func NewCreateProductUseCase(productRepository entity.ProductsRepository) *CreateProductUseCase {
	return &CreateProductUseCase{ProductRepository: productRepository}
}

// o use case não necessita saber como será executado apenas precisa ter declarado
// a estrutura de entrada
// a estrutura de saida
// e o que o usecase irá utilizar

func (u *CreateProductUseCase) Execute(input CreateProductInputDto) (*CreateProductOutputDto, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return &CreateProductOutputDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
