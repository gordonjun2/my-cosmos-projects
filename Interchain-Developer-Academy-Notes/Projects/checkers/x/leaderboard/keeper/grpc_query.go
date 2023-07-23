package keeper

import (
	"github.com/alice/checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
