package otherworld_test

import (
	"testing"

	keepertest "github.com/b9lab/other-world/testutil/keeper"
	"github.com/b9lab/other-world/testutil/nullify"
	"github.com/b9lab/other-world/x/otherworld"
	"github.com/b9lab/other-world/x/otherworld/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WorldParams: &types.WorldParams{
			Name:       "55",
			Gravity:    78,
			LandSupply: 23,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OtherworldKeeper(t)
	otherworld.InitGenesis(ctx, *k, genesisState)
	got := otherworld.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.WorldParams, got.WorldParams)
	// this line is used by starport scaffolding # genesis/test/assert
}
