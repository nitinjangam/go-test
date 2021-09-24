package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(t *testing.T) {
	t.Run("test correct path", func(t *testing.T) {
		m := mux.NewRouter()
		m.HandleFunc("/", homeHandler)
		m.HandleFunc("/products", productsHandler)
		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			log.Fatalf("Error while creating request %v", err)
		}
		w := httptest.NewRecorder()
		m.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("Want %v, Got %v", 200, w.Code)
		}
		r, err = http.NewRequest(http.MethodGet, "/products", nil)
		if err != nil {
			log.Fatalf("Error while creating request %v", err)
		}
		w = httptest.NewRecorder()
		m.ServeHTTP(w, r)
		if w.Code != 200 {
			t.Errorf("Want %v, Got %v", 200, w.Code)
		}
	})
}

func TestWrongPath(t *testing.T) {
	t.Run("test wrong path", func(t *testing.T) {
		m := mux.NewRouter()
		m.HandleFunc("/", homeHandler)
		m.HandleFunc("/products", productsHandler)
		r, err := http.NewRequest(http.MethodGet, "/nit", nil)
		if err != nil {
			log.Fatalf("Error while creating request %v", err)
		}
		w := httptest.NewRecorder()
		m.ServeHTTP(w, r)
		if w.Code != 404 {
			t.Errorf("Want %v, Got %v", 404, w.Code)
		}
	})
}
