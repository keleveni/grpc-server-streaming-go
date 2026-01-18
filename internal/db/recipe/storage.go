package recipe

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	"github.com/uptrace/bun"
)

type Storage struct {
	db *bun.DB
}

func NewStorage(db *bun.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Create(ctx context.Context, recipe *recipe.Recipe) error {
	_, err := s.db.NewInsert().Model(recipe).Exec(ctx)

	return err
}

func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*recipe.Recipe, error) {
	entity := recipe.Recipe{ID: id}
	if err := s.db.NewSelect().Model(&entity).WherePK().Scan(ctx); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(recipe.ErrRecipeNotFound)
		}

		return nil, err
	}

	return &entity, nil
}
