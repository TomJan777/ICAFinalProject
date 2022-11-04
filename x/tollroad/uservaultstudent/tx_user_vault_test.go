package uservaultstudent_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/b9lab/toll-road/testutil/network"
	"github.com/b9lab/toll-road/x/tollroad/client/cli"
	"github.com/b9lab/toll-road/x/tollroad/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize
var moduleAddress = sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName))).String()

func TestCreateUserVault(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"111"}
	moduleBalanceQueryArgs := []string{
		moduleAddress,
		fmt.Sprintf("--%s=%s", sdkcli.FlagDenom, net.Config.BondDenom),
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	for _, tc := range []struct {
		desc                string
		idRoadOperatorIndex string
		idToken             string

		args []string
		err  error
		code uint32
	}{
		{
			idRoadOperatorIndex: strconv.Itoa(0),
			idToken:             net.Config.BondDenom,

			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idRoadOperatorIndex,
				tc.idToken,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUserVault(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)

				if tc.code == 0 {
					out, err = clitestutil.ExecTestCLICmd(ctx, sdkcli.GetBalancesCmd(), moduleBalanceQueryArgs)
					require.NoError(t, err)
					require.Equal(t,
						fmt.Sprintf("{\"denom\":\"stake\",\"amount\":\"%s\"}\n", "111"),
						out.String())
				}
			}
		})
	}
}

func TestUpdateUserVault(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"111"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
	args := []string{
		"0",
		net.Config.BondDenom,
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUserVault(), args)
	require.NoError(t, err)
	moduleBalanceQueryArgs := []string{
		moduleAddress,
		fmt.Sprintf("--%s=%s", sdkcli.FlagDenom, net.Config.BondDenom),
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	for _, tc := range []struct {
		desc                string
		idRoadOperatorIndex string
		idToken             string

		fields []string
		args   []string
		code   uint32
		err    error
	}{
		{
			desc:                "valid no change",
			idRoadOperatorIndex: strconv.Itoa(0),
			idToken:             net.Config.BondDenom,

			fields: fields,
			args:   common,
		},
		{
			desc:                "valid increase vault",
			idRoadOperatorIndex: strconv.Itoa(0),
			idToken:             net.Config.BondDenom,

			fields: []string{"112"},
			args:   common,
		},
		{
			desc:                "valid decrease vault",
			idRoadOperatorIndex: strconv.Itoa(0),
			idToken:             net.Config.BondDenom,

			fields: []string{"110"},
			args:   common,
		},
		{
			desc:                "key not found",
			idRoadOperatorIndex: strconv.Itoa(100000),
			idToken:             net.Config.BondDenom,

			fields: []string{"112"},
			args:   common,
			code:   sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idRoadOperatorIndex,
				tc.idToken,
			}
			args = append(args, tc.fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateUserVault(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)

				if tc.code == 0 {
					out, err = clitestutil.ExecTestCLICmd(ctx, sdkcli.GetBalancesCmd(), moduleBalanceQueryArgs)
					require.NoError(t, err)
					require.Equal(t,
						fmt.Sprintf("{\"denom\":\"stake\",\"amount\":\"%s\"}\n", tc.fields[0]),
						out.String())
				}
			}
		})
	}
}

func TestDeleteUserVault(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"111"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
	args := []string{
		"0",
		net.Config.BondDenom,
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUserVault(), args)
	require.NoError(t, err)
	moduleBalanceQueryArgs := []string{
		moduleAddress,
		fmt.Sprintf("--%s=%s", sdkcli.FlagDenom, net.Config.BondDenom),
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	for _, tc := range []struct {
		desc                string
		idRoadOperatorIndex string
		idToken             string

		args []string
		code uint32
		err  error
	}{
		{
			desc:                "valid",
			idRoadOperatorIndex: strconv.Itoa(0),
			idToken:             net.Config.BondDenom,

			args: common,
		},
		{
			desc:                "key not found",
			idRoadOperatorIndex: strconv.Itoa(100000),
			idToken:             net.Config.BondDenom,

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idRoadOperatorIndex,
				tc.idToken,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteUserVault(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				var resp sdk.TxResponse
				require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, tc.code, resp.Code)

				if tc.code == 0 {
					out, err = clitestutil.ExecTestCLICmd(ctx, sdkcli.GetBalancesCmd(), moduleBalanceQueryArgs)
					require.NoError(t, err)
					require.Equal(t,
						"{\"denom\":\"stake\",\"amount\":\"0\"}\n",
						out.String())
				}
			}
		})
	}
}
