package cli

import (
	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateRoadOperator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-road-operator [index] [name] [token] [active]",
		Short: "Create a new RoadOperator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get value arguments
			argName := args[0]
			argToken := args[1]
			argActive, err := cast.ToBoolE(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRoadOperator(
				clientCtx.GetFromAddress().String(),
				argName,
				argToken,
				argActive,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRoadOperator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-road-operator [index] [name] [token] [active]",
		Short: "Update a RoadOperator",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argName := args[1]
			argToken := args[2]
			argActive, err := cast.ToBoolE(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRoadOperator(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argName,
				argToken,
				argActive,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteRoadOperator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-road-operator [index]",
		Short: "Delete a RoadOperator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteRoadOperator(
				clientCtx.GetFromAddress().String(),
				indexIndex,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
