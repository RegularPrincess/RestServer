// DBHendler
package interfaces

import "database/sql"

type DBHendler interface {
	Execute(query string) error
	Query(query string) (sql.Rows, error)
	Create(e Entity) error
	Close()
}
