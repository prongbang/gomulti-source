package driver

import "github.com/jmoiron/sqlx"

// SQLxDriver is the interface
type SQLxDriver interface {
	Connect() *sqlx.DB
}
