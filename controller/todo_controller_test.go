package controller

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/koga456/sample-api/controller/dto"
	"github.com/koga456/sample-api/test"
)

func TestGetTodos_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.GetTodos(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse)
	if len(todosResponse.Todos) != 0 {
		t.Errorf("Response is %v", todosResponse.Todos)
	}
}

func TestGetTodos_ExistTodo(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryGetTodosExist{})
	target.GetTodos(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse.Todos)
	if len(todosResponse.Todos) != 2 {
		t.Errorf("Response is %v", todosResponse.Todos)
	}
}

func TestGetTodos_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.GetTodos(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	if w.Body.Len() != 0 {
		t.Errorf("body is %v", w.Body.Len())
	}
}

func TestPostTodo_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todos/", json)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PostTodo(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != r.Host+r.URL.Path+"2" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPostTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/todos/", json)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.PostTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != "" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPutTodo_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/2", json)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/2", json)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.PutTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_OK(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/2", nil)

	target := NewTodoController(&test.MockTodoRepository{})
	target.DeleteTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.DeleteTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteTodo_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/2", nil)

	target := NewTodoController(&test.MockTodoRepositoryError{})
	target.DeleteTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}
