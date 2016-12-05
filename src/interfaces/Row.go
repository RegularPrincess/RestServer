// Row
package interfaces

type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
}
