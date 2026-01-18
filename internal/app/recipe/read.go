package recipe

import (
	"context"

	"github.com/google/uuid"
)

type ReadUseCase struct {
	storage Storage
}

func NewReadUseCase(
	storage Storage,
) *ReadUseCase {
	return &ReadUseCase{
		storage: storage,
	}
}

func (u *ReadUseCase) Get(ctx context.Context, id string) (*Recipe, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	recipe, err := u.storage.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}
