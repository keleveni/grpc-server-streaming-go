package recipe

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
	"github.com/kodaikumatani/grpc-golang/internal/db/follow"
	"github.com/kodaikumatani/grpc-golang/internal/orm"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList_Followee(t *testing.T) {
	ctx := t.Context()

	db := lo.Must(orm.NewBun(ctx))
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	_, err = tx.NewTruncateTable().Model((*recipe.Recipe)(nil)).Cascade().Exec(ctx)
	require.NoError(t, err)

	_, err = tx.NewTruncateTable().Model((*follow.Follow)(nil)).Cascade().Exec(ctx)
	require.NoError(t, err)

	followInfo := follow.Follow{
		FollowerID: "01HZ0000000000000000000001",
		FolloweeID: "01HZ0000000000000000000002",
	}

	// Create recipes
	recipes := []recipe.Recipe{
		{
			ID:     lo.Must(uuid.NewV7()),
			UserID: followInfo.FolloweeID,
		},
		{
			ID:     lo.Must(uuid.NewV7()),
			UserID: "01HZ0000000000000000000003",
		},
	}

	_, err = tx.NewInsert().Model(&recipes).Exec(ctx)
	require.NoError(t, err)

	// Create follow relationship
	_, err = tx.NewInsert().Model(&followInfo).Exec(ctx)
	require.NoError(t, err)

	// Test: List recipes from followees
	actual, err := NewStorage(tx).List().
		Followee(followInfo.FollowerID).
		Exec(ctx)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, followInfo.FolloweeID, actual[0].UserID)
}

func TestList_CreatedAfter(t *testing.T) {
	ctx := t.Context()

	db := lo.Must(orm.NewBun(ctx))
	tx, err := db.Begin()
	require.NoError(t, err)
	defer tx.Rollback()

	_, err = tx.NewTruncateTable().Model((*recipe.Recipe)(nil)).Cascade().Exec(ctx)
	require.NoError(t, err)

	_, err = tx.NewTruncateTable().Table("recipes").Exec(ctx)
	require.NoError(t, err)

	// Setup test data
	baseTime := time.Now().Add(-1 * time.Hour)

	records := []recipe.Recipe{
		{
			ID:        lo.Must(uuid.NewV7()),
			CreatedAt: baseTime.Add(-1 * time.Hour),
			UpdatedAt: baseTime.Add(-1 * time.Hour),
		},
		{
			ID:        lo.Must(uuid.NewV7()),
			CreatedAt: baseTime.Add(30 * time.Minute),
			UpdatedAt: baseTime.Add(30 * time.Minute),
		},
	}

	_, err = tx.NewInsert().Model(&records).Exec(ctx)
	require.NoError(t, err)

	// Test: List recipes created after base time
	actual, err := NewStorage(tx).List().
		CreatedAfter(baseTime).
		Exec(ctx)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, records[1].ID, actual[0].ID)
}
