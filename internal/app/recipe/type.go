package recipe

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

var (
	ErrRecipeNotFound = errors.New("recipe not found")
)

type Recipe struct {
	ID          uuid.UUID `bun:"id,pk"`
	UserID      string    `bun:"user_id"`
	Title       string    `bun:"title"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}

type Storage interface {
	Create(ctx context.Context, recipe *Recipe) error
	GetByID(ctx context.Context, id uuid.UUID) (*Recipe, error)
}
