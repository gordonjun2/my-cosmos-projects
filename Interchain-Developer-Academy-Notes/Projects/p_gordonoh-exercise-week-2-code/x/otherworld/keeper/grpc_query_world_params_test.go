package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/b9lab/other-world/testutil/keeper"
	"github.com/b9lab/other-world/testutil/nullify"
	"github.com/b9lab/other-world/x/otherworld/types"
)

func TestWorldParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.OtherworldKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestWorldParams(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWorldParamsRequest
		response *types.QueryGetWorldParamsResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetWorldParamsRequest{},
			response: &types.QueryGetWorldParamsResponse{WorldParams: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WorldParams(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
