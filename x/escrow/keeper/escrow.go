package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ethermint/x/escrow/types"
)

func (k Keeper) EscrowCoin(ctx sdk.Context, buyer sdk.AccAddress, amount sdk.Coins) error {

	// TODO: Support only 1 coin
	if len(amount) != 1 {
		return sdkerrors.Wrapf(types.ErrOnlyOneDenomAllowed, "%d denoms included", len(amount))
	}
	prefix := "hale"
	if !strings.HasPrefix(amount[0].Denom, prefix) {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom was: %s", amount[0].Denom)
	}
	// Escrow source tokens. It fails if balance insufficient.
	escrowAddress := types.GetEscrowAddress()
	return k.BankKeeper.SendCoins(ctx, buyer, escrowAddress, amount)
}

func (k Keeper) GetDayReceiverAmount(ctx sdk.Context, blockHeight int64, receiver string) int64 {
	storeKey := types.DayReceiverAmountStoreKey(blockHeight, receiver)
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
