package keeper_test

import (
	"testing"

	testkeeper "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/leaderboard/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeaderboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
