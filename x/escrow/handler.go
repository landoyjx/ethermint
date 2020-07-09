package escrow

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/ethermint/x/escrow/types"
)

// NewHandler creates the msg handler of this module, as required by Cosmos-SDK standard.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		// types.Logger.Info(fmt.Sprintf("msg: %+v", msg))
		switch msg := msg.(type) {
		case MsgSendWithUnlock:
			return handlePlaceBet(ctx, msg, keeper)
		case MsgPayout:
			return handlePayout(ctx, msg, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handlePlaceBet(ctx sdk.Context, msg MsgSendWithUnlock, keeper Keeper) (*sdk.Result, error) {
	types.Logger.Info("MsgSendWithUnlock")

	if err := keeper.EscrowCoin(ctx, msg.FromAddress, msg.Amount); err != nil {
		return nil, err
	}

	amount := msg.Amount[0].Amount.Uint64()
	receiver := msg.ToAddress.String()

	types.Logger.Info(fmt.Sprintf("Updating mappings due to receiver(%s) amount(%d)", receiver, amount))

	DayId := types.GetDayId(ctx.BlockTime().Unix())

	totalAmount := keeper.GetDayReceiverAmount(ctx, DayId, receiver)
	keeper.SetDayReceiverAmount(ctx, DayId, receiver, totalAmount+int64(amount))

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handlePayout(ctx sdk.Context, msg MsgPayout, keeper Keeper) (*sdk.Result, error) {
	types.Logger.Info("Payout request")

	receiver := msg.Receiver.String()

	if keeper.GetDayReceiverPaid(ctx, msg.DayId, receiver) {
		return nil, sdkerrors.Wrapf(types.Error, "receiver(%s) already paid in day (%d)", receiver, msg.DayId)
	}
	amount := keeper.GetAmountByDayId(ctx, msg.DayId, receiver)
	if amount == 0 {
		return nil, sdkerrors.Wrapf(types.Error, "receiver(%s) has no coin in the day(%d)", receiver, msg.DayId)
	}
	keeper.SetDayReceiverPaid(ctx, msg.DayId, receiver, true)
	keeper.Payout(ctx, msg.Receiver, sdk.NewCoins(sdk.NewInt64Coin("hale", amount)))

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
