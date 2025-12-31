package recipe

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

var (
	ErrRecipeNotFound = errors.New("recipe not found")
)

type Recipe struct {
	ID          uuid.UUID `bun:"id,pk"`
	UserID      ulid.ULID `bun:"user_id"`
	Title       string    `bun:"title"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
