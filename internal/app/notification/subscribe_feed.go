package notification

import (
	"context"
	"time"

	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
)

type SubscribeFeedUseCase struct {
	recipeStorage recipe.Storage
}

func NewSubscribeFeedUseCase(
	recipeStorage recipe.Storage,
) *SubscribeFeedUseCase {
	return &SubscribeFeedUseCase{
		recipeStorage: recipeStorage,
	}
}

func (u *SubscribeFeedUseCase) StreamRecipes(
	ctx context.Context,
	userID string,
	callback func(recipe.Recipe) error,
) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	lastCheck := time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Get recent recipes from followed users
			recipes, err := u.recipeStorage.
				List().
				Followee(userID).
				CreatedAfter(lastCheck).
				Exec(ctx)
			if err != nil {
				return err
			}

			// Send each recipe through callback
			for _, r := range recipes {
				if err := callback(r); err != nil {
					return err
				}
			}

			lastCheck = time.Now()
		}
	}
}
