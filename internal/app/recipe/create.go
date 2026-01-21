package recipe

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type CreateUseCase struct {
	storage Storage
}

func NewCreateUseCase(
	storage Storage,
) *CreateUseCase {
	return &CreateUseCase{
		storage: storage,
	}
}

func (u *CreateUseCase) Invoke(
	ctx context.Context,
	userID,
	title,
	description string,
) (*Recipe, error) {
	recipe := Recipe{
		ID:          lo.Must(uuid.NewV7()),
		UserID:      userID,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := u.storage.Create(ctx, &recipe); err != nil {
		return nil, err
	}

	return &recipe, nil
}
