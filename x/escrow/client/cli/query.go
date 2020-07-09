package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/ethermint/x/escrow/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinPriceBetCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the coin_price_bet module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	coinPriceBetCmd.AddCommand(flags.GetCommands(
	//GetCmdTodayCoinPrices(storeKey, cdc),
	)...)

	return coinPriceBetCmd
}
