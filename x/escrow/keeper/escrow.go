package keeper

import (
	"encoding/json"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ethermint/x/escrow/types"
)

func (k Keeper) EscrowCoin(ctx sdk.Context, buyer sdk.AccAddress, amount sdk.Coins) error {

	// TODO: Support only 1 coin
	if len(amount) != 1 {
		return sdkerrors.Wrapf(types.ErrOnlyOneDenomAllowed, "%d denoms included", len(amount))
	}
	prefix := "uhale"
	if !strings.HasPrefix(amount[0].Denom, prefix) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom was: %s", amount[0].Denom)
	}
	// Escrow source tokens. It fails if balance insufficient.
	escrowAddress := types.GetEscrowAddress()
	return k.BankKeeper.SendCoins(ctx, buyer, escrowAddress, amount)
}

type DayIdInfo struct {
	DayId      int64     `json:"day_id"`
	UnlockTime time.Time `json:"unlock_time"`
	Amount     uint64    `json:"amount"`
}

func (k Keeper) NewDayIdInfo(dayId int64, unlockTime time.Time, amount uint64) (res []DayIdInfo) {
	info := DayIdInfo{
		DayId:      dayId,
		UnlockTime: unlockTime,
		Amount:     amount,
	}

	return append(res, info)
}

func (k Keeper) GetReceiverDayIdsInfo(ctx sdk.Context, receiver string) (res []DayIdInfo) {
	storeKey := types.ReceiverStoreKey(receiver)
	store := ctx.KVStore(k.storeKey)
	if b := store.Get(storeKey); b != nil {

		if err := json.Unmarshal(b, &res); err != nil {
			return
		} else {
			return
		}

	} else {
		return
	}
}

func (k Keeper) SetReceiverDayIdsInfo(ctx sdk.Context, receiver string, dayInfos []DayIdInfo) {
	storeKey := types.ReceiverStoreKey(receiver)
	store := ctx.KVStore(k.storeKey)

	jsonBytes, _ := json.Marshal(dayInfos)
	store.Set(storeKey, jsonBytes)
}

func (k Keeper) GetDayReceiverAmount(ctx sdk.Context, dayId int64, receiver string) int64 {
	storeKey := types.DayReceiverAmountStoreKey(dayId, receiver)
	store := ctx.KVStore(k.storeKey)
	if b := store.Get(storeKey); b != nil {
		return types.BytesToInt64(b)
	} else {
		return 0
	}
}

func (k Keeper) SetDayReceiverAmount(ctx sdk.Context, dayId int64, receiver string, amount int64) {
	storeKey := types.DayReceiverAmountStoreKey(dayId, receiver)
	store := ctx.KVStore(k.storeKey)
	store.Set(storeKey, types.Int64ToBytes(amount))
}

func (k Keeper) GetDayReceiverPaid(ctx sdk.Context, dayId int64, receiver string) bool {
	storeKey := types.DayReceiverPaidStoreKey(dayId, receiver)
	store := ctx.KVStore(k.storeKey)
	var paid bool
	if b := store.Get(storeKey); b != nil {
		k.cdc.MustUnmarshalBinaryBare(b, &paid)
	}
	return paid
}

func (k Keeper) SetDayReceiverPaid(ctx sdk.Context, dayId int64, receiver string, paid bool) {
	storeKey := types.DayReceiverPaidStoreKey(dayId, receiver)
	store := ctx.KVStore(k.storeKey)
	store.Set(storeKey, k.cdc.MustMarshalBinaryBare(paid))
}

func (k Keeper) GetAmountByDayId(ctx sdk.Context, dayId int64, receiver string) int64 {
	return k.GetDayReceiverAmount(ctx, dayId, receiver)
}

func (k Keeper) Payout(ctx sdk.Context, receiver sdk.AccAddress, amount sdk.Coins) {
	escrowAddress := types.GetEscrowAddress()
	k.BankKeeper.SendCoins(ctx, escrowAddress, receiver, amount)
}
