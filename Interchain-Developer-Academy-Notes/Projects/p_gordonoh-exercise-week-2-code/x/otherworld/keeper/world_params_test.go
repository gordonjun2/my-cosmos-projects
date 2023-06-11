package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/other-world/testutil/keeper"
	"github.com/b9lab/other-world/testutil/nullify"
	"github.com/b9lab/other-world/x/otherworld/keeper"
	"github.com/b9lab/other-world/x/otherworld/types"
)

func createTestWorldParams(keeper *keeper.Keeper, ctx sdk.Context) types.WorldParams {
	item := types.WorldParams{}
	keeper.SetWorldParams(ctx, item)
	return item
}

func TestWorldParamsGet(t *testing.T) {
	keeper, ctx := keepertest.OtherworldKeeper(t)
	item := createTestWorldParams(keeper, ctx)
	rst, found := keeper.GetWorldParams(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestWorldParamsRemove(t *testing.T) {
	keeper, ctx := keepertest.OtherworldKeeper(t)
	createTestWorldParams(keeper, ctx)
	keeper.RemoveWorldParams(ctx)
	_, found := keeper.GetWorldParams(ctx)
	require.False(t, found)
}
