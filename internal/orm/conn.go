package orm

import (
	"context"
	"database/sql"
	"os"

	"github.com/cockroachdb/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewBun(ctx context.Context) (*bun.DB, error) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:25432"),
		pgdriver.WithTLSConfig(nil),
		pgdriver.WithUser(os.Getenv("POSTGRES_USER")),
		pgdriver.WithPassword(os.Getenv("POSTGRES_PASSWORD")),
		pgdriver.WithDatabase(os.Getenv("POSTGRES_DATABASE")),
	)

	db := bun.NewDB(sql.OpenDB(pgconn), pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "could not ping the database")
	}

	return db, nil
}
