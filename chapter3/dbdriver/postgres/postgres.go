package postgres

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

// bean
type PostgresDriver struct{}

// connect
func (dr PostgresDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *PostgresDriver

// init is called prior to main.
func init() {
	d = new(PostgresDriver)
	sql.Register("postgres", d)
}
