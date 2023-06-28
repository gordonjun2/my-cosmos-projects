package keeper

import (
	"context"
	"strconv"

	rules "github.com/b9lab/checkers/x/checkers/rules"
	"github.com/b9lab/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextGame, found := k.Keeper.GetNextGame(ctx)
	if !found {
		panic("NextGame not found")
	}
	newIndex := strconv.FormatUint(nextGame.IdValue, 10)
	newGame := rules.New()
	storedGame := types.StoredGame{
		Creator:   msg.Creator,
		Index:     newIndex,
		Game:      newGame.String(),
		Turn:      rules.PieceStrings[newGame.Turn],
		Red:       msg.Red,
		Black:     msg.Black,
		MoveCount: 0,
		BeforeId:  types.NoFifoIdKey,
		AfterId:   types.NoFifoIdKey,
		Deadline:  types.FormatDeadline(types.GetNextDeadline(ctx)),
		Winner:    rules.PieceStrings[rules.NO_PLAYER],
		Wager:     msg.Wager,
		Token:     msg.Token,
	}
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SendToFifoTail(ctx, &storedGame, &nextGame)
	k.Keeper.SetStoredGame(ctx, storedGame)

	nextGame.IdValue++
	k.Keeper.SetNextGame(ctx, nextGame)

	ctx.GasMeter().ConsumeGas(types.CreateGameGas, "Create game")

	// What to emit
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.StoredGameEventKey),
			sdk.NewAttribute(types.StoredGameEventCreator, msg.Creator),
			sdk.NewAttribute(types.StoredGameEventIndex, newIndex),
			sdk.NewAttribute(types.StoredGameEventRed, msg.Red),
			sdk.NewAttribute(types.StoredGameEventBlack, msg.Black),
			sdk.NewAttribute(types.StoredGameEventWager, strconv.FormatUint(msg.Wager, 10)),
			sdk.NewAttribute(types.StoredGameEventToken, msg.Token),
		),
	)

	return &types.MsgCreateGameResponse{
		IdValue: newIndex,
	}, nil
}
