// @title API de Conversão Ilumeo
// @version 1.0
// @description API que calcula a taxa de conversão por minuto com filtros
// @host localhost:8080
// @BasePath /

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "ilumeo/docs"
	"ilumeo/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Swagger em /swagger/index.html
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Rota principal da API
	http.HandleFunc("/api/conversao", handlers.HandleConversao)

	fmt.Println("API ouvindo em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
