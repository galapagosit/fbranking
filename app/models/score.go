package models

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

func RegisterScore(Txn *sql.Tx, id string, score int, scoreBest int) {
	var _user_id string
	var lastScore int
	err := Txn.QueryRow("SELECT user_id, score FROM user_score WHERE user_id=?", id).Scan(&_user_id, &lastScore)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user_score with id:", id)

		// 新規レコード作成
		query := "INSERT INTO user_score (user_id, score, created_at) values(?, ?, ?)"
		_, err := Txn.Exec(query, id, scoreBest, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatal("insert error: ", err)
		}
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("user exist id:", id)

		if scoreBest > lastScore {
			query := "UPDATE user_score SET score=? WHERE user_id=?"
			_, err := Txn.Exec(query, scoreBest, id)
			if err != nil {
				log.Fatal("insert error: ", err)
			}
		}
	}

	logging(Txn, id, score, scoreBest)
}

func logging(Txn *sql.Tx, id string, score int, scoreBest int) {
	query := "INSERT INTO user_score_log (user_id, score, score_best, created_at) values(?, ?, ?, ?)"
	_, err := Txn.Exec(query, id, score, scoreBest, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatal("insert error: ", err)
	}
}

func GetScores(Txn *sql.Tx, idList []string) []interface{} {
	var idListParam []interface{} = make([]interface{}, len(idList))
	for i, d := range idList {
		idListParam[i] = d
	}

	rows, err := Txn.Query("SELECT user_id, score FROM user_score WHERE user_id IN (?"+strings.Repeat(",?", len(idList)-1)+")", idListParam...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	res := []interface{}{}
	for rows.Next() {
		var user_id string
		var score int
		if err := rows.Scan(&user_id, &score); err != nil {
			log.Fatal(err)
		}
		res = append(res, map[string]interface{}{"id": user_id, "score": score})
	}
	return res
}
