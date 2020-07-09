package escrow

import (
	"github.com/cosmos/ethermint/x/escrow/keeper"
	"github.com/cosmos/ethermint/x/escrow/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier
	logger        = types.Logger
)

type (
	Keeper            = keeper.Keeper
	MsgSendWithUnlock = types.MsgSendWithUnlock
	MsgPayout         = types.MsgPayout
)
