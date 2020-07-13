package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ethermint/x/escrow/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryReceiver = "receiver"
)

// NewQuerier is the module level router for state queries.
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		types.Logger.Info(fmt.Sprintf("query /%s", path[0]))
		switch path[0] {
		case QueryReceiver:
			return queryReceiverInfo(ctx, keeper, req, path[1])
		}
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
	}
}
