package recipe

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kodaikumatani/grpc-golang/internal/app/recipe"
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

	_, err = tx.NewTruncateTable().Table("recipes").Exec(ctx)
	require.NoError(t, err)

	// Setup test data
	followerID := "01HZ0000000000000000000001"
	followeeID := "01HZ0000000000000000000002"
	otherUserID := "01HZ0000000000000000000003"

	// Create recipes
	records := []recipe.Recipe{
		{
			ID:     lo.Must(uuid.NewV7()),
			UserID: followeeID,
		},
		{
			ID:     lo.Must(uuid.NewV7()),
			UserID: otherUserID,
		},
	}

	_, err = tx.NewInsert().Model(&records).Exec(ctx)
	require.NoError(t, err)

	// Create follow relationship
	_, err = tx.Exec(
		"INSERT INTO follows (follower_id, followee_id, created_at) VALUES (?, ?, ?)",
		followerID, followeeID, time.Time{})
	require.NoError(t, err)

	// Test: List recipes from followees
	actual, err := NewStorage(tx).List().
		Followee(followerID).
		Exec(ctx)
	require.NoError(t, err)

	assert.Len(t, actual, 1)
	assert.Equal(t, followeeID, actual[0].UserID)
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
