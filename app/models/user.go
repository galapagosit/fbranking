package models

import (
	"database/sql"
	"log"
	"time"
)

func CreateUser(Txn *sql.Tx, id string) {
	var _id int
	err := Txn.QueryRow("SELECT id FROM user WHERE id=?", id).Scan(&_id)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id:", id)

		// 新規ユーザ作成
		query := "INSERT INTO user (id, created_at) values(?, ?)"
		_, err := Txn.Exec(query, id, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatal("insert error: ", err)
		}
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("user exist id:", id)
	}
}
