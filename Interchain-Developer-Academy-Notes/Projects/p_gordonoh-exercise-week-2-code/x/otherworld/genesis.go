package otherworld

import (
	"github.com/b9lab/other-world/x/otherworld/keeper"
	"github.com/b9lab/other-world/x/otherworld/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.WorldParams != nil {
		k.SetWorldParams(ctx, *genState.WorldParams)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all worldParams
	worldParams, found := k.GetWorldParams(ctx)
	if found {
		genesis.WorldParams = &worldParams
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
