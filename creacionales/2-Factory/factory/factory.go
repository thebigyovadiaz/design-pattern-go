package factory

import (
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/2-Factory/connection"
)

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.Postgres{}
	default:
		return nil
	}
}
