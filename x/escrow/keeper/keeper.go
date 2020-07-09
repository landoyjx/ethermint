package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/ethermint/x/escrow/types"
)

// Keeper of the scavenge store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	BankKeeper types.BankKeeper
}

// NewKeeper creates a escrow keeper
func NewKeeper(cdc *codec.Codec, coinKeeper bank.Keeper, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		BankKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
