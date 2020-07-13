package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/ethermint/x/escrow/keeper"
	"github.com/cosmos/ethermint/x/escrow/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	coinPriceBetCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the escrow module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	coinPriceBetCmd.AddCommand(flags.GetCommands(
		GetCmdReceiverEscrow(storeKey, cdc),
	)...)

	return coinPriceBetCmd
}

func GetCmdReceiverEscrow(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  keeper.QueryReceiver,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			receiver := args[0]
			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/%s/%s", queryRoute, keeper.QueryReceiver,receiver),
				nil,
			)
			if err != nil {
				fmt.Printf("read request fail -  %s\n", err)
				return nil
			}

			dayInfos := []keeper.DayIdInfo{}

			cdc.MustUnmarshalJSON(res, &dayInfos)

			return cliCtx.PrintOutput(dayInfos)
		},
	}
}
