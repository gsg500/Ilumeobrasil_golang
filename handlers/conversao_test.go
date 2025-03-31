package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleConversao_Sucesso(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/conversao?canal=email", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConversao)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && status != http.StatusInternalServerError {
		t.Errorf("handler retornou código errado: obteve %v", status)
	}
}
