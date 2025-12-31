package recipe

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ReadUseCase struct {
	db bun.IDB
}

func NewReadUseCase(
	db bun.IDB,
) *ReadUseCase {
	return &ReadUseCase{
		db: db,
	}
}

func (u *ReadUseCase) Get(ctx context.Context, id string) (*Recipe, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	recipe := Recipe{ID: uid}
	if err := u.db.NewSelect().Model(&recipe).WherePK().Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(ErrRecipeNotFound)
		}

		return nil, err
	}

	return &recipe, nil
}
