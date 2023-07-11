package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/leaderboard/testutil/keeper"
	"github.com/cosmonaut/leaderboard/x/leaderboard/keeper"
	"github.com/cosmonaut/leaderboard/x/leaderboard/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
