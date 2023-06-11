package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/b9lab/other-world/testutil/network"
	"github.com/b9lab/other-world/testutil/nullify"
	"github.com/b9lab/other-world/x/otherworld/client/cli"
	"github.com/b9lab/other-world/x/otherworld/types"
)

func networkWithWorldParamsObjects(t *testing.T) (*network.Network, types.WorldParams) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	worldParams := &types.WorldParams{}
	nullify.Fill(&worldParams)
	state.WorldParams = worldParams
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.WorldParams
}

func TestShowWorldParams(t *testing.T) {
	net, obj := networkWithWorldParamsObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.WorldParams
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowWorldParams(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetWorldParamsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.WorldParams)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.WorldParams),
				)
			}
		})
	}
}
