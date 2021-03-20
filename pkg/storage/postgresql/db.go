package postgresql

import (
	"context"
	"database/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//DB ...
type DB struct {
	pgsql *sql.DB
}

//NewDB ..
func NewDB(db *sql.DB) *DB {
	return &DB{pgsql: db}
}

// Connect returns SQL database connection.
func (db *DB) Connect(ctx context.Context) (*sql.Conn, error) {
	c,err := db.pgsql.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to connect to database : "+err.Error())
	}
	return c, nil
}