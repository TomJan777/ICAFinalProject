package uservaultstudent_test

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/toll-road/testutil/keeper"
	"github.com/b9lab/toll-road/testutil/mock_types"
	"github.com/b9lab/toll-road/x/tollroad"
	"github.com/b9lab/toll-road/x/tollroad/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setupMsgServerWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *mock_types.MockBankEscrowKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := mock_types.NewMockBankEscrowKeeper(ctrl)
	k, ctx := keepertest.TollroadKeeperWithMocks(t, bankMock)
	tollroad.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, bankMock
}

func TestUserVaultMsgServerCreateFive(t *testing.T) {
	srv, k, wctx, ctrl, escrow := setupMsgServerWithMock(t)
	ctx := sdk.UnwrapSDKContext(wctx)

	defer ctrl.Finish()

	creator := "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	for i := 0; i < 5; i++ {
		createRequest := &types.MsgCreateUserVault{
			Creator:           creator,
			RoadOperatorIndex: strconv.Itoa(rand.Intn(1000)),
			Token:             sdk.DefaultBondDenom,
			Balance:           uint64(rand.Intn(1000)),
		}
		escrow.ExpectFundVault(wctx, creator, sdk.DefaultBondDenom, createRequest.Balance)
		_, err := srv.CreateUserVault(wctx, createRequest)
		require.NoError(t, err)
		rst, found := k.GetUserVault(ctx,
			createRequest.Creator,
			createRequest.RoadOperatorIndex,
			createRequest.Token,
		)
		require.True(t, found)
		require.EqualValues(t, types.UserVault{
			Owner:             creator,
			RoadOperatorIndex: createRequest.RoadOperatorIndex,
			Token:             createRequest.Token,
			Balance:           createRequest.Balance,
		}, rst)
	}
}

func TestUserVaultMsgServerCreateExists(t *testing.T) {
	srv, _, wctx, ctrl, escrow := setupMsgServerWithMock(t)

	defer ctrl.Finish()

	creator := "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	for i := 0; i < 5; i++ {
		createRequest := &types.MsgCreateUserVault{
			Creator:           creator,
			RoadOperatorIndex: strconv.Itoa(rand.Intn(1000)),
			Token:             sdk.DefaultBondDenom,
			Balance:           uint64(rand.Intn(1000)),
		}
		escrow.ExpectFundVault(wctx, creator, sdk.DefaultBondDenom, createRequest.Balance)
		srv.CreateUserVault(wctx, createRequest)
		_, err := srv.CreateUserVault(wctx, createRequest)
		require.ErrorIs(t, err, types.ErrIndexSet)
	}
}

func TestUserVaultMsgServerCreateCases(t *testing.T) {
	creator := "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"

	for _, tc := range []struct {
		desc       string
		request    *types.MsgCreateUserVault
		expectBank bool
		bankErr    error
		err        error
	}{
		{
			desc: "ErrorZero",
			request: &types.MsgCreateUserVault{
				Creator:           creator,
				RoadOperatorIndex: strconv.Itoa(rand.Intn(1000)),
				Token:             sdk.DefaultBondDenom,
				Balance:           0,
			},
			err: types.ErrZeroTokens,
		},
		{
			desc: "ErrorBank",
			request: &types.MsgCreateUserVault{
				Creator:           creator,
				RoadOperatorIndex: strconv.Itoa(rand.Intn(1000)),
				Token:             sdk.DefaultBondDenom,
				Balance:           100,
			},
			expectBank: true,
			bankErr:    errors.New("bank error"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, k, wctx, ctrl, escrow := setupMsgServerWithMock(t)
			ctx := sdk.UnwrapSDKContext(wctx)

			defer ctrl.Finish()

			if tc.expectBank {
				call := escrow.ExpectFundVault(wctx, creator, tc.request.Token, tc.request.Balance)
				if tc.bankErr != nil {
					call.Return(tc.bankErr)
				}
			}

			_, err := srv.CreateUserVault(wctx, tc.request)

			if tc.bankErr != nil {
				require.ErrorIs(t, err, tc.bankErr)
			} else if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetUserVault(ctx,
					tc.request.Creator,
					tc.request.RoadOperatorIndex,
					tc.request.Token,
				)
				require.True(t, found)
				require.EqualValues(t, types.UserVault{
					Owner:             tc.request.Creator,
					RoadOperatorIndex: tc.request.RoadOperatorIndex,
					Token:             tc.request.Token,
					Balance:           tc.request.Balance,
				}, rst)
			}
		})
	}
}

func TestUserVaultMsgServerUpdate(t *testing.T) {
	creator := "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	roadOperatorIndex := strconv.Itoa(rand.Intn(1000))
	balance := uint64(500)

	for _, tc := range []struct {
		desc       string
		request    *types.MsgUpdateUserVault
		expectBank bool
		bankErr    error
		err        error
		panic      string
	}{
		{
			desc: "CompletedAndIncreased",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           1000,
			},
			expectBank: true,
		},
		{
			desc: "CompletedAndDecreased",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           400,
			},
			expectBank: true,
		},
		{
			desc: "IncreaseFailed",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           1000,
			},
			expectBank: true,
			bankErr:    errors.New("bank error"),
		},
		{
			desc: "RefundFailed",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           400,
			},
			expectBank: true,
			bankErr:    errors.New("bank error"),
			panic:      "bank error",
		},
		{
			desc: "CannotToZero",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           0,
			},
			err: types.ErrZeroTokens,
		},
		{
			desc: "KeyNotFoundByOwner",
			request: &types.MsgUpdateUserVault{
				Creator:           "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g",
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           1000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "KeyNotFoundByRoadOperatorIndex",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: strconv.Itoa(100000),
				Token:             sdk.DefaultBondDenom,
				Balance:           1000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "KeyNotFoundByToken",
			request: &types.MsgUpdateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             "not-stake",
				Balance:           1000,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, k, wctx, ctrl, escrow := setupMsgServerWithMock(t)
			ctx := sdk.UnwrapSDKContext(wctx)

			defer ctrl.Finish()

			escrow.ExpectFundVault(wctx, creator, sdk.DefaultBondDenom, balance)

			createRequest := &types.MsgCreateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           balance,
			}
			_, err := srv.CreateUserVault(wctx, createRequest)
			require.NoError(t, err)

			if tc.expectBank {
				var call *gomock.Call
				if tc.request.Balance < balance {
					call = escrow.ExpectRefundVault(wctx, creator, sdk.DefaultBondDenom, balance-tc.request.Balance)
				} else if balance < tc.request.Balance {
					call = escrow.ExpectFundVault(wctx, creator, sdk.DefaultBondDenom, tc.request.Balance-balance)
				}
				if tc.bankErr != nil {
					call.Return(tc.bankErr)
				}
			}

			if tc.panic != "" {
				defer func() {
					r := recover()
					require.NotNil(t, r, "The code did not panic")
					require.Equal(t, tc.panic, r)
				}()
			}

			_, err = srv.UpdateUserVault(wctx, tc.request)
			if tc.panic == "" {
				if tc.bankErr != nil {
					require.ErrorIs(t, err, tc.bankErr)
				} else if tc.err != nil {
					require.ErrorIs(t, err, tc.err)
				} else {
					require.NoError(t, err)
					rst, found := k.GetUserVault(ctx,
						createRequest.Creator,
						createRequest.RoadOperatorIndex,
						createRequest.Token,
					)
					require.True(t, found)
					require.EqualValues(t, types.UserVault{
						Owner:             creator,
						RoadOperatorIndex: roadOperatorIndex,
						Token:             sdk.DefaultBondDenom,
						Balance:           tc.request.Balance,
					}, rst)
				}
			}
		})
	}
}

func TestUserVaultMsgServerDelete(t *testing.T) {
	creator := "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	roadOperatorIndex := strconv.Itoa(rand.Intn(1000))

	for _, tc := range []struct {
		desc       string
		request    *types.MsgDeleteUserVault
		expectBank bool
		bankErr    error
		err        error
		panic      string
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
			},
			expectBank: true,
		},
		{
			desc: "RefundFailed",
			request: &types.MsgDeleteUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
			},
			expectBank: true,
			bankErr:    errors.New("bank error"),
			panic:      "bank error",
		},
		{
			desc: "KeyNotFoundByOwner",
			request: &types.MsgDeleteUserVault{
				Creator:           "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g",
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "KeyNotFoundByRoadOperatorIndex",
			request: &types.MsgDeleteUserVault{
				Creator:           creator,
				RoadOperatorIndex: strconv.Itoa(100000),
				Token:             sdk.DefaultBondDenom,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "KeyNotFoundByToken",
			request: &types.MsgDeleteUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             "not-stake",
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, k, wctx, ctrl, escrow := setupMsgServerWithMock(t)
			ctx := sdk.UnwrapSDKContext(wctx)

			defer ctrl.Finish()

			escrow.ExpectFundVault(wctx, creator, sdk.DefaultBondDenom, 500)

			_, err := srv.CreateUserVault(wctx, &types.MsgCreateUserVault{
				Creator:           creator,
				RoadOperatorIndex: roadOperatorIndex,
				Token:             sdk.DefaultBondDenom,
				Balance:           500,
			})
			require.NoError(t, err)

			if tc.expectBank {
				call := escrow.ExpectRefundVault(wctx, creator, sdk.DefaultBondDenom, 500)
				if tc.bankErr != nil {
					call.Return(tc.bankErr)
				}
			}

			if tc.panic != "" {
				defer func() {
					r := recover()
					require.NotNil(t, r, "The code did not panic")
					require.Equal(t, tc.panic, r)
				}()
			}

			_, err = srv.DeleteUserVault(wctx, tc.request)

			if tc.panic == "" {
				if tc.bankErr != nil {
					require.ErrorIs(t, err, tc.bankErr)
				} else if tc.err != nil {
					require.ErrorIs(t, err, tc.err)
				} else {
					require.NoError(t, err)
					_, found := k.GetUserVault(ctx,
						tc.request.Creator,
						tc.request.RoadOperatorIndex,
						tc.request.Token,
					)
					require.False(t, found)
				}
			}
		})
	}
}
