package keeper

import (

	// "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func queryReceiverInfo(
	ctx sdk.Context, keeper Keeper, req abci.RequestQuery, receiver string,
) ([]byte, error) {

	dayInfos := keeper.GetReceiverDayIdsInfo(ctx, receiver)

	return keeper.cdc.MustMarshalJSON(dayInfos), nil
}
