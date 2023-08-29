package models

import (
	"main/database"
)

func Get(id int64) (todo Todo, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()
	rows := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)
	err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.Done)

	return

}
