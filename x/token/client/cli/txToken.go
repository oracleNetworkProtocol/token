package cli

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/oracleNetworkProtocol/token/x/token/types"
)

func CmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [symbol] [originalSymbol] [desc] [wholeName] [totalSupply] [own] [mintable]",
		Short: "Create a new Token",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSymbol, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsOriginalSymbol, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsDesc, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsWholeName, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsTotalSupply, err := cast.ToInt64E(args[4])
			if err != nil {
				return err
			}
			argsOwn, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsMintable, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}
			cmd.Flags().Set(flags.FlagFrom, argsOwn)
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			log.Println("create address:", argsOwn)
			log.Println("create address bech32 []byte:", clientCtx.GetFromAddress())
			msg := types.NewMsgCreateToken(clientCtx.GetFromAddress().String(), argsSymbol, argsOriginalSymbol, argsDesc, argsWholeName, argsTotalSupply, argsOwn, argsMintable)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token [id] [symbol] [originalSymbol] [desc] [wholeName] [totalSupply] [own] [mintable]",
		Short: "Update a Token",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsSymbol, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsOriginalSymbol, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsDesc, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsWholeName, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			argsTotalSupply, err := cast.ToInt64E(args[5])
			if err != nil {
				return err
			}

			argsOwn, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}

			argsMintable, err := cast.ToBoolE(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateToken(clientCtx.GetFromAddress().String(), id, argsSymbol, argsOriginalSymbol, argsDesc, argsWholeName, argsTotalSupply, argsOwn, argsMintable)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-token [id]",
		Short: "Delete a Token by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteToken(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
