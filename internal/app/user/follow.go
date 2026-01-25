package user

import (
	"context"
)

type FollowUseCase struct {
	storage FollowStorage
}

func NewFollowUseCase(
	storage FollowStorage,
) *FollowUseCase {
	return &FollowUseCase{
		storage: storage,
	}
}

func (u *FollowUseCase) Follow(
	ctx context.Context,
	followerID, followeeID string,
) error {
	return u.storage.Follow(ctx, followerID, followeeID)
}
