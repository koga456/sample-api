package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/koga456/sample-api/model/entity"
)

type TodoRepository interface {
	GetTodos() (todos []entity.TodoEntity, err error)
	InsertTodo(todo entity.TodoEntity) (id int, err error)
	UpdateTodo(todo entity.TodoEntity) (err error)
	DeleteTodo(id int) (err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []entity.TodoEntity, err error) {
	todos = []entity.TodoEntity{}
	rows, err := Db.
		Query("SELECT id, title, content FROM todo ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		todo := entity.TodoEntity{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			log.Print(err)
			return
		}
		todos = append(todos, todo)
	}

	return
}

func (tr *todoRepository) InsertTodo(todo entity.TodoEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Print(err)
		return
	}
	err = Db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}

func (tr *todoRepository) UpdateTodo(todo entity.TodoEntity) (err error) {
	_, err = Db.Exec("UPDATE todo SET title = ?, content = ? WHERE id = ?", todo.Title, todo.Content, todo.Id)
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	_, err = Db.Exec("DELETE FROM todo WHERE id = ?", id)
	return
}
