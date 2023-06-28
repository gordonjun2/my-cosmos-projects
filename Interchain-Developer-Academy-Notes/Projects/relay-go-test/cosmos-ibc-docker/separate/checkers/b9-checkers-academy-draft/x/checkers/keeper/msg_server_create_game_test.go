package keeper_test

import (
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *IntegrationTestSuite) TestCreateGame() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	createResponse, err := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   12,
		Token:   sdk.DefaultBondDenom,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(types.MsgCreateGameResponse{
		IdValue: "1",
	}, *createResponse)
}

func (suite *IntegrationTestSuite) TestCreateGameDidNotPay() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   12,
		Token:   sdk.DefaultBondDenom,
	})
	suite.RequireBankBalance(balAlice, alice)
	suite.RequireBankBalance(balBob, bob)
	suite.RequireBankBalance(balCarol, carol)
	suite.RequireBankBalance(0, checkersModuleAddress)
}

func (suite *IntegrationTestSuite) TestCreate1GameHasSaved() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   13,
		Token:   sdk.DefaultBondDenom,
	})
	keeper := suite.app.CheckersKeeper
	nextGame, found := keeper.GetNextGame(suite.ctx)
	suite.Require().True(found)
	suite.Require().EqualValues(types.NextGame{
		Creator:  "",
		IdValue:  2,
		FifoHead: "1",
		FifoTail: "1",
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
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     13,
		Token:     "stake",
	}, game1)
}

func (suite *IntegrationTestSuite) TestCreate1GameGetAll() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   14,
		Token:   sdk.DefaultBondDenom,
	})
	keeper := suite.app.CheckersKeeper
	games := keeper.GetAllStoredGame(suite.ctx)
	suite.Require().Len(games, 1)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     14,
		Token:     "stake",
	}, games[0])
}

func (suite *IntegrationTestSuite) TestCreate1GameEmitted() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   15,
		Token:   sdk.DefaultBondDenom,
	})
	events := sdk.StringifyEvents(suite.ctx.EventManager().ABCIEvents())
	suite.Require().Len(events, 1)
	event := events[0]
	suite.Require().EqualValues(sdk.StringEvent{
		Type: "message",
		Attributes: []sdk.Attribute{
			{Key: "module", Value: "checkers"},
			{Key: "action", Value: "NewGameCreated"},
			{Key: "Creator", Value: alice},
			{Key: "Index", Value: "1"},
			{Key: "Red", Value: bob},
			{Key: "Black", Value: carol},
			{Key: "Wager", Value: "15"},
			{Key: "Token", Value: "stake"},
		},
	}, event)
}

func (suite *IntegrationTestSuite) TestCreate1GameConsumedGas() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	gasBefore := suite.ctx.GasMeter().GasConsumed()
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   15,
		Token:   sdk.DefaultBondDenom,
	})
	gasAfter := suite.ctx.GasMeter().GasConsumed()
	suite.Require().Equal(uint64(13_400+10), gasAfter-gasBefore)
}

func (suite *IntegrationTestSuite) TestCreateGameRedAddressBad() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	createResponse, err := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     "notanaddress",
		Black:   carol,
		Wager:   16,
		Token:   sdk.DefaultBondDenom,
	})
	suite.Require().Nil(createResponse)
	suite.Require().Equal(
		"red address is invalid: notanaddress: decoding bech32 failed: invalid index of 1",
		err.Error())
}

func (suite *IntegrationTestSuite) TestCreateGameEmptyRedAddress() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	createResponse, err := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     "",
		Black:   carol,
		Wager:   16,
		Token:   sdk.DefaultBondDenom,
	})
	suite.Require().Nil(createResponse)
	suite.Require().Equal(
		"red address is invalid: : empty address string is not allowed",
		err.Error())
}

func (suite *IntegrationTestSuite) TestCreate3Games() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   17,
		Token:   sdk.DefaultBondDenom,
	})
	createResponse2, err2 := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   18,
		Token:   sdk.DefaultBondDenom,
	})
	suite.Require().Nil(err2)
	suite.Require().EqualValues(types.MsgCreateGameResponse{
		IdValue: "2",
	}, *createResponse2)
	createResponse3, err3 := suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   19,
		Token:   sdk.DefaultBondDenom,
	})
	suite.Require().Nil(err3)
	suite.Require().EqualValues(types.MsgCreateGameResponse{
		IdValue: "3",
	}, *createResponse3)
}

func (suite *IntegrationTestSuite) TestCreate3GamesHasSaved() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   20,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   21,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   22,
		Token:   sdk.DefaultBondDenom,
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
		AfterId:   "2",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     20,
		Token:     "stake",
	}, game1)
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
		BeforeId:  "1",
		AfterId:   "3",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     21,
		Token:     "stake",
	}, game2)
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
		BeforeId:  "2",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     22,
		Token:     "stake",
	}, game3)
}

func (suite *IntegrationTestSuite) TestCreate3GamesGetAll() {
	suite.setupSuiteWithBalances()
	goCtx := sdk.WrapSDKContext(suite.ctx)

	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
		Wager:   23,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: bob,
		Red:     carol,
		Black:   alice,
		Wager:   24,
		Token:   sdk.DefaultBondDenom,
	})
	suite.msgServer.CreateGame(goCtx, &types.MsgCreateGame{
		Creator: carol,
		Red:     alice,
		Black:   bob,
		Wager:   25,
		Token:   sdk.DefaultBondDenom,
	})
	keeper := suite.app.CheckersKeeper
	games := keeper.GetAllStoredGame(suite.ctx)
	suite.Require().Len(games, 3)
	suite.Require().EqualValues(types.StoredGame{
		Creator:   alice,
		Index:     "1",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       bob,
		Black:     carol,
		MoveCount: uint64(0),
		BeforeId:  "-1",
		AfterId:   "2",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     23,
		Token:     "stake",
	}, games[0])
	suite.Require().EqualValues(types.StoredGame{
		Creator:   bob,
		Index:     "2",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       carol,
		Black:     alice,
		MoveCount: uint64(0),
		BeforeId:  "1",
		AfterId:   "3",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     24,
		Token:     "stake",
	}, games[1])
	suite.Require().EqualValues(types.StoredGame{
		Creator:   carol,
		Index:     "3",
		Game:      "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:      "b",
		Red:       alice,
		Black:     bob,
		MoveCount: uint64(0),
		BeforeId:  "2",
		AfterId:   "-1",
		Deadline:  types.FormatDeadline(suite.ctx.BlockTime().Add(types.MaxTurnDuration)),
		Winner:    "*",
		Wager:     25,
		Token:     "stake",
	}, games[2])
}
