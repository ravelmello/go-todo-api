package models

import "main/database"

func Delete(id int64) (int64, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	delete := `DELETE FROM todos WHERE id = $1`
	response, err := conn.Exec(delete, id)

	if err != nil {
		return 0, err
	}

	return response.RowsAffected()
}
