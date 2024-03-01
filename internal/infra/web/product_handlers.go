package web

import (
	"encoding/json"
	"net/http"

	usecase "github.com/Patrikesm/kafka-with-go/internal/usecase/product"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductUseCase *usecase.ListProductUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductUseCase:   listProductUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	//apesar de ser um controller ele não sabe o que está sendo feito
	//dentro da aplicação ou seja toda regra de negócio ainda está em apenas um lugar
	var input usecase.CreateProductInputDto

	//hidrate this datas to fit into DTO structure
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//aqui estou definindo o content-type em caso de go puro poderia ser um middleware
	//além de enviar os dados e encodar a estrutura em json novamente
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
