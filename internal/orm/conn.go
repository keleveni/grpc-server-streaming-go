package orm

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewBun(ctx context.Context) (*bun.DB, error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:25432"),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser("gcook"),
		pgdriver.WithPassword("p4ssw0rd!"),
		pgdriver.WithDatabase("gcook"),
	)

	db := bun.NewDB(sql.OpenDB(pgconn), pgdialect.New())

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "could not ping the database")
	}

	return db, nil
}
