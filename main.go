// @title API de Conversão Ilumeo
// @version 1.0
// @description API que calcula a taxa de conversão por minuto com filtros
// @host localhost:8080
// @BasePath /

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "ilumeo/docs"
	"ilumeo/handlers"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Swagger em /swagger/index.html
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Rota principal da API com middleware CORS
	http.HandleFunc("/api/conversao", corsMiddleware(handlers.HandleConversao))

	fmt.Println("🚀 API ouvindo em :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
