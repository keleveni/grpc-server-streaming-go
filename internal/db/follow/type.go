package follow

import (
	"time"

	"github.com/uptrace/bun"
)

type Follow struct {
	bun.BaseModel `bun:"table:follows"`

	FollowerID string    `bun:"follower_id"`
	FolloweeID string    `bun:"followee_id"`
	CreatedAt  time.Time `bun:"created_at"`
}
