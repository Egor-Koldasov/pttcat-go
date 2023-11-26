package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

func init() {
	var err error
	dbConn, err = sql.Open("postgres", "user=admin dbname=pttcat password=devpassword sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func SaveState(state string) {
	_, err := dbConn.Exec(
		"INSERT INTO State(userId, dataJson, updatedAt, createdAt) VALUES('1', $1, NOW(), NOW()) ON CONFLICT (userId) DO UPDATE SET dataJson = $1, updatedAt = NOW()",
		state,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func GetState() string {
	var dataJson string
	err := dbConn.QueryRow("SELECT dataJson FROM State WHERE userId = '1'").Scan(&dataJson)
	if err != nil {
		log.Fatal(err)
	}
	return dataJson
}
