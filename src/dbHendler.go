// DBHendler
package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	I "github.com/regularprincess/RestServer/src/interfaces"
)

func NewPostgreDB(connInfo string) (DB, error) {
	conn, err := sql.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	db := new(DB)
	db.conn = conn
	return db, nil
}

type DB struct {
	conn *sql.DB
}

func (db *DB) Execute(query string) error {
	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) Query(query string) (I.Row, error) {
	rows, err := db.conn.Query(query)
	if err != nil {
		return err, nil
	}
	myRows := new(MyRow)
	myRows.rows = rows
	return myRows
}

type MyRow struct {
	rows *sql.Rows
}

func (r MyRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r MyRow) Next() bool {
	return r.Rows.Next()
}
