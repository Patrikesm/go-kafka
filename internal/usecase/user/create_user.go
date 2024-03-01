package usecase

import "github.com/Patrikesm/kafka-with-go/internal/entity"

type CreateUserInputDto struct {
	Name string
}

type CreateUserOutputDto struct {
	ID   string
	Name string
}

type CreateUseCase struct {
	UserRepository entity.UserRepository
}

func NewCreateUseCase(userRepository entity.UserRepository) *CreateUseCase {
	return &CreateUseCase{UserRepository: userRepository}
}

func (c *CreateUseCase) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	user := entity.NewUser(input.Name)

	err := c.UserRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDto{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
