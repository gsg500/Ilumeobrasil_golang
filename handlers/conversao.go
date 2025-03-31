package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Resultado struct {
	Intervalo     time.Time `json:"intervalo"`       // Timestamp agrupado por minuto
	Total         int       `json:"total_registros"` // Total de registros por intervalo
	Conversoes    int       `json:"conversoes"`      // Total de conversões no intervalo
	TaxaConversao float64   `json:"taxa_conversao"`  // Taxa de conversão calculada
}

// HandleConversao godoc
// @Summary Retorna evolução da taxa de conversão
// @Description Agrupamento por minuto com possibilidade de filtro por canal e intervalo de data
// @Tags Conversão
// @Accept json
// @Produce json
// @Param canal query string false "Canal de origem (ex: email, MOBILE)"
// @Param inicio query string false "Data de início (formato ISO)"
// @Param fim query string false "Data de fim (formato ISO)"
// @Success 200 {array} Resultado
// @Failure 500 {string} string "Erro interno"
// @Router /api/conversao [get]

func HandleConversao(w http.ResponseWriter, r *http.Request) {
	db, err := setupDB()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar no banco: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	canal := r.URL.Query().Get("canal")
	inicio := r.URL.Query().Get("inicio")
	fim := r.URL.Query().Get("fim")

	baseQuery := `
		SELECT
			date_trunc('minute', created_at) AS intervalo,
			COUNT(*) AS total_registros,
			COUNT(CASE WHEN response_status_id = 1 THEN 1 END) AS conversoes,
			ROUND((COUNT(CASE WHEN response_status_id = 1 THEN 1 END)::decimal / COUNT(*)::decimal)*100, 2) AS taxa_conversao
		FROM inside.users_surveys_responses_aux
		WHERE 1=1
	`

	var args []interface{}
	argIdx := 1

	if canal != "" {
		baseQuery += fmt.Sprintf(" AND origin = $%d", argIdx)
		args = append(args, canal)
		argIdx++
	}
	if inicio != "" {
		baseQuery += fmt.Sprintf(" AND created_at >= $%d", argIdx)
		args = append(args, inicio)
		argIdx++
	}
	if fim != "" {
		baseQuery += fmt.Sprintf(" AND created_at <= $%d", argIdx)
		args = append(args, fim)
		argIdx++
	}

	baseQuery += " GROUP BY intervalo ORDER BY intervalo"

	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro na consulta: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var resultados []Resultado
	for rows.Next() {
		var r Resultado
		if err := rows.Scan(&r.Intervalo, &r.Total, &r.Conversoes, &r.TaxaConversao); err != nil {
			http.Error(w, "Erro ao ler resultado", http.StatusInternalServerError)
			return
		}
		resultados = append(resultados, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

func setupDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
