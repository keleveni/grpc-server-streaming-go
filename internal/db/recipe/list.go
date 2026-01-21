package recipe

import (
	"context"
	"time"

	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	"github.com/uptrace/bun"
)

type List struct {
	query *bun.SelectQuery
}

func (l *List) Followee(userID string) recipe.List {
	l.query = l.query.
		Join("JOIN follows ON follows.followee_id = recipe.user_id").
		Where("follows.follower_id = ?", userID)
	return l
}

func (l *List) CreatedAfter(t time.Time) recipe.List {
	l.query = l.query.Where("recipe.created_at > ?", t)
	return l
}

func (l *List) Exec(ctx context.Context) ([]recipe.Recipe, error) {
	var recipes []recipe.Recipe
	if err := l.query.
		Model(&recipes).
		Scan(ctx); err != nil {
		return nil, err
	}

	return recipes, nil
}
