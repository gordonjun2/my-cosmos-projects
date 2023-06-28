package checkers

import (
	"github.com/b9lab/checkers/x/checkers/keeper"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set if defined
	if genState.Leaderboard != nil {
		k.SetLeaderboard(ctx, *genState.Leaderboard)
	}

	// Set all the playerInfo
	for _, elem := range genState.PlayerInfoList {
		k.SetPlayerInfo(ctx, *elem)
	}

	// Set all the storedGame
	for _, elem := range genState.StoredGameList {
		k.SetStoredGame(ctx, *elem)
	}

	// Set if defined
	if genState.NextGame != nil {
		k.SetNextGame(ctx, *genState.NextGame)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all leaderboard
	leaderboard, found := k.GetLeaderboard(ctx)
	if found {
		genesis.Leaderboard = &leaderboard
	}

	// Get all playerInfo
	playerInfoList := k.GetAllPlayerInfo(ctx)
	for _, elem := range playerInfoList {
		elem := elem
		genesis.PlayerInfoList = append(genesis.PlayerInfoList, &elem)
	}

	// Get all storedGame
	storedGameList := k.GetAllStoredGame(ctx)
	for _, elem := range storedGameList {
		elem := elem
		genesis.StoredGameList = append(genesis.StoredGameList, &elem)
	}

	// Get all nextGame
	nextGame, found := k.GetNextGame(ctx)
	if found {
		genesis.NextGame = &nextGame
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
