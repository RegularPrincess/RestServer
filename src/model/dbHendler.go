// DBHendler
package model

import (
	"database/sql"

	_ "github.com/lib/pq"
	I "github.com/regularprincess/RestServer/src/interfaces"
)

func NewPostgreDB(connInfo string) (DB, error) {
	conn, err := sql.Open("postgres", connInfo)
	db := new(DB)
	if err != nil {
		return *db, err
	}
	err = conn.Ping()
	if err != nil {
		return *db, err
	}
	db.conn = conn
	return *db, nil
}

type DB struct {
	conn *sql.DB
}

func (db *DB) Execute(query string) error {
	_, err := db.conn.Exec(query)
	return err
}
func (db *DB) Create(e I.Entity) (int, error) {
	str := "INSERT INTO " + e.GetTableName() + "(" + e.GetTableFields() + ") VALUES (" + e.GetFields() + ") RETURNING id;"
	//fmt.Print(str)
	row, err := db.Query(str)
	defer row.Close()

	if err != nil {
		return 0, err
	}
	var id int
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (db *DB) Query(query string) (sql.Rows, error) {
	rows, err := db.conn.Query(query)
	if err != nil {
		return *rows, nil
	}
	return *rows, nil
}
