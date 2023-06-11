package keeper

import (
	"context"

	"github.com/b9lab/other-world/x/otherworld/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WorldParams(c context.Context, req *types.QueryGetWorldParamsRequest) (*types.QueryGetWorldParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetWorldParams(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWorldParamsResponse{WorldParams: val}, nil
}
