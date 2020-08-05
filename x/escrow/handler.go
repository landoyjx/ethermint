package escrow

import (
	"fmt"
	"time"

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
			return handleSendWithUnlock(ctx, msg, keeper)
		case MsgPayout:
			return handlePayout(ctx, msg, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleSendWithUnlock(ctx sdk.Context, msg MsgSendWithUnlock, keeper Keeper) (*sdk.Result, error) {

	types.Logger.Info("MsgSendWithUnlock")

	dayId := types.GetDayId(msg.UnlockTime.Unix())

	if err := keeper.EscrowCoin(ctx, msg.FromAddress, msg.Amount); err != nil {
		return nil, err
	}

	blockTimeDayId := types.GetDayId(ctx.BlockTime().Unix())

	if dayId < blockTimeDayId+1 {
		return nil, sdkerrors.Wrapf(types.Error, "unlock time: %v block time: %v escrow earliest time: %v,dayId: %v blockTimeDayId+1: %v",
			msg.UnlockTime, ctx.BlockTime(), time.Unix(types.GetTimeByDayId(blockTimeDayId+1), 0).UTC(), dayId, blockTimeDayId+1)
	}

	amount := msg.Amount[0].Amount.Uint64()
	receiver := msg.ToAddress.String()

	types.Logger.Info(fmt.Sprintf("Updating mappings due to receiver(%s) amount(%d)", receiver, amount))

	totalAmount := keeper.GetDayReceiverAmount(ctx, dayId, receiver)
	keeper.SetDayReceiverAmount(ctx, dayId, receiver, totalAmount+int64(amount))

	if oldDaysInfo := keeper.GetReceiverDayIdsInfo(ctx, receiver); len(oldDaysInfo) == 0 {
		keeper.SetReceiverDayIdsInfo(ctx, receiver, keeper.NewDayIdInfo(dayId, msg.UnlockTime, amount))
	} else {
		exist := false
		for k, v := range oldDaysInfo {
			if v.DayId == dayId {
				v.Amount = v.Amount + amount
				oldDaysInfo[k] = v
				keeper.SetReceiverDayIdsInfo(ctx, receiver, oldDaysInfo)
				exist = true
				break
			}
		}

		if !exist {
			keeper.SetReceiverDayIdsInfo(ctx, receiver, append(oldDaysInfo, keeper.NewDayIdInfo(dayId, msg.UnlockTime, amount)...))
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handlePayout(ctx sdk.Context, msg MsgPayout, keeper Keeper) (*sdk.Result, error) {
	types.Logger.Info("Payout request")

	blockDayId := types.GetDayId(ctx.BlockTime().Unix())

	if blockDayId <= msg.DayId {
		return nil, sdkerrors.Wrapf(types.Error, "blockDayId < = msg.DayId,invalid dayId(%d) unlock time: %v, current blockDayId(%d) block time: %v",
			msg.DayId, time.Unix(types.GetTimeByDayId(msg.DayId), 0).UTC(), blockDayId, ctx.BlockTime())
	}
	receiver := msg.Receiver.String()

	if keeper.GetDayReceiverPaid(ctx, msg.DayId, receiver) {
		return nil, sdkerrors.Wrapf(types.Error, "receiver(%s) already paid in day (%d)", receiver, msg.DayId)
	}
	amount := keeper.GetAmountByDayId(ctx, msg.DayId, receiver)
	if amount == 0 {
		return nil, sdkerrors.Wrapf(types.Error, "receiver(%s) has no coin in the day(%d)", receiver, msg.DayId)
	}
	keeper.SetDayReceiverPaid(ctx, msg.DayId, receiver, true)
	keeper.Payout(ctx, msg.Receiver, sdk.NewCoins(sdk.NewInt64Coin("uhale", amount)))

	daysInfo := keeper.GetReceiverDayIdsInfo(ctx, msg.Receiver.String())
	for k, v := range daysInfo {
		if v.DayId == msg.DayId {
			daysInfo = append(daysInfo[:k], daysInfo[k+1:]...)
			keeper.SetReceiverDayIdsInfo(ctx, receiver, daysInfo)
			break
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
