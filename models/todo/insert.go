package models

import (
	"main/database"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()
	insert := `INSERT INTO todos (title, description, createdAt, updatedAt, isDone) values ($1, $2, $3, $4, $5) returning id`
	err = conn.QueryRow(insert, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt, todo.Done).Scan(&id)
	return
}
