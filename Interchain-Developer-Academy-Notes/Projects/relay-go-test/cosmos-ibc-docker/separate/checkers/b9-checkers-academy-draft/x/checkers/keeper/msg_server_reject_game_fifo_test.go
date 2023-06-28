package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestRejectSecondGameHasSavedFifo() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "1",
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  3,
		FifoHead: "2",
		FifoTail: "2",
	}, nextGame)
	game2, found2 := keeper.GetStoredGame(suite.ctx, "2")
	suite.Require().True(found2)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     12,
		Token:     "stake",
	}, game2)
}

func (suite *IntegrationTestSuite) TestRejectMiddleGameHasSavedFifo() {
	suite.setupSuiteWithOneGameForRejectGame()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   12,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   13,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.RejectGame(goCtx, &types.MsgRejectGame{
		Creator: carol,
		IdValue: "2",
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  4,
		FifoHead: "1",
		FifoTail: "3",
	}, nextGame)
	game1, found1 := keeper.GetStoredGame(suite.ctx, "1")
	suite.Require().True(found1)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "3",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     11,
		Token:     "stake",
	}, game1)
	game3, found3 := keeper.GetStoredGame(suite.ctx, "3")
	suite.Require().True(found3)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   carol,
		Index:     "3",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       alice,
		Black:     bob,
		MoveCount: uint64(0),
		BeforeId:  "1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     13,
		Token:     "stake",
	}, game3)
}
