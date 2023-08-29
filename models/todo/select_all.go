package models

import "main/database"

func GetAll() (todos []Todo, err error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()
	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.Done)
		if err != nil {
			continue //TODO: make log error here
		}
		todos = append(todos, todo)
	}

	return

}
