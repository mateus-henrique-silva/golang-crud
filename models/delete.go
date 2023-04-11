package models

import "go/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`Delete FROM todos WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
