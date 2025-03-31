# Ilumeo - Desafio Tech Lead 🚀

API desenvolvida para mostrar a **evolução temporal da taxa de conversão** por minuto, com filtros por canal, data de início e fim.

---

## 📦 Tecnologias

- Golang 1.22
- PostgreSQL
- Docker & Docker Compose
- Swagger (documentação automática)
- godotenv

---

## 🧪 Endpoints

### `GET /api/conversao`

Consulta a taxa de conversão por minuto.

#### Parâmetros (query):

- `canal` (string) — canal de origem (ex: `MOBILE`)
- `inicio` (string - ISO 8601) — data inicial
- `fim` (string - ISO 8601) — data final

---

## 🔧 Variáveis de Ambiente (`.env`)

```env
DB_HOST=seu_host
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=nome_do_banco

```
---
## ⚙️ Otimização de Performance (Banco de Dados)

Para melhorar o desempenho das consultas, é altamente recomendada a criação de um índice composto nos campos utilizados nos filtros e agrupamento:

```sql
CREATE INDEX idx_created_at_origin 
ON inside.users_surveys_responses_aux (created_at, origin);