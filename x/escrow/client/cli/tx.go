package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/ethermint/x/escrow/types"
	"github.com/spf13/cobra"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Escrow coin transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		SendTxCmd(cdc),
		PayoutTxCmd(cdc),
	)
	return txCmd
}

func SendTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [from_key_or_address] [to_address] [amount] [unlock_time]",
		Short: "Create and/or sign and broadcast a MsgSendWithUnlock transaction",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInputAndFrom(inBuf, args[0]).WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			unlockTime, err := time.ParseInLocation("2006-01-02 15:04:05", args[3], time.UTC)
			if err != nil {
				return err
			}

			tomorrow := time.Now().Add(10 * time.Minute)

			fmt.Printf("unlock time: %v  tomorrow: %v \n", unlockTime, tomorrow.UTC())

			// if unlockTime.Before(tomorrow) {
			// 	return fmt.Errorf("unlock time: %v before tomorrow: %v ", unlockTime, tomorrow)
			// }

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgSendWithUnlock(cliCtx.GetFromAddress(), to, coins, unlockTime)
			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = flags.PostCommands(cmd)[0]

	return cmd
}

func PayoutTxCmd(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "payout [from_key_or_address] [receiver_address] [day_id]",
		Short: "Create and/or sign and broadcast a payout transaction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInputAndFrom(inBuf, args[0]).WithCodec(cdc)

			receiver, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			dayId, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgPayout(receiver, dayId)

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd = flags.PostCommands(cmd)[0]

	return cmd

}
