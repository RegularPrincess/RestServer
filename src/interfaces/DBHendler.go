// DBHendler
package interfaces

type DBHendler interface {
	Execute(query string)
	Query(query string) Row
}
