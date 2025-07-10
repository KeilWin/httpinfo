package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFaviconHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/favicon.ico", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(FaviconHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler return unexpected status code: got %v, want %v", status, http.StatusNotFound)
	}

	expectedBody := "Not found\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("Handler return unexpected body:\ngot %v\nwant %v", rr.Body.String(), expectedBody)
	}
}
