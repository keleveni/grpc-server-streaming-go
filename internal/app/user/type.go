package user

import "context"

type FollowStorage interface {
	Follow(ctx context.Context, followerID, followeeID string) error
}
