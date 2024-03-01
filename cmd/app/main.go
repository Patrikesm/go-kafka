package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Patrikesm/kafka-with-go/internal/infra/web"

	"github.com/Patrikesm/kafka-with-go/internal/infra/akafka"
	"github.com/Patrikesm/kafka-with-go/internal/infra/repository"
	usecase "github.com/Patrikesm/kafka-with-go/internal/usecase/product"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repo := repository.NewProductRepository(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repo)
	listProductUseCase := usecase.NewListProductUseCase(repo)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductHandler)

	go http.ListenAndServe(":8080", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products", "users"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}

		//convert json to struct
		//cleaning all informations to defined dto
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		//Insert into database
		_, err = createProductUseCase.Execute(dto)
	}
}
