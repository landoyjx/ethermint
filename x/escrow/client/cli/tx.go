package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/cosmos/cosmos-sdk/version"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	//authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/ethermint/x/escrow/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinPriceBetCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Coin Price Bet transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	coinPriceBetCmd.AddCommand(flags.PostCommands(
	// GetCmdSetChannel(cdc),
	// GetCmdPlaceBet(cdc),
	)...)

	return coinPriceBetCmd
}
