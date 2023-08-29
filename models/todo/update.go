package models

import "main/database"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	update := `UPDATE todos SET title = $1, description = $2, createdAt = $3, updatedAt = $4, done = $5 WHERE id = 6$`
	response, err := conn.Exec(update, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
