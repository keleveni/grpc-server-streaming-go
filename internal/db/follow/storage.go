package follow

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Storage struct {
	db bun.IDB
}

func NewStorage(db bun.IDB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Follow(
	ctx context.Context,
	followerID, followeeID string,
) error {
	model := Follow{
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  time.Now(),
	}
	_, err := s.db.NewInsert().Model(&model).Exec(ctx)

	return err
}
