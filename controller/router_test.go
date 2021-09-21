package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/koga456/sample-api/test"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := NewRouter(&test.MockTodoController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/todos/", target.HandleTodosRequest)

}

func TestGetTodos(t *testing.T) {
	r, _ := http.NewRequest("GET", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	r, _ := http.NewRequest("POST", "/todos/", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	r, _ := http.NewRequest("PUT", "/todos/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/todos/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 405 {
		t.Errorf("Response cod is %v", w.Code)
	}
}
