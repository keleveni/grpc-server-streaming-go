package recipe

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

type CreateUseCase struct {
	db bun.IDB
}

func NewCreateUseCase(
	db bun.IDB,
) *CreateUseCase {
	return &CreateUseCase{
		db: db,
	}
}

func (u *CreateUseCase) Invoke(
	ctx context.Context,
	userID,
	title,
	description string,
) (*Recipe, error) {
	if _, err := ulid.Parse(userID); err != nil {
		return nil, err
	}

	recipe := Recipe{
		ID:          lo.Must(uuid.NewV7()),
		UserID:      userID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if _, err := u.db.NewInsert().Model(&recipe).Exec(ctx); err != nil {
		return nil, err
	}

	return &recipe, nil
}
