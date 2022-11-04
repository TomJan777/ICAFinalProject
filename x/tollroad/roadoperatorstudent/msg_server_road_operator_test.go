package roadoperatorstudent_test

import (
	"math/rand"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/b9lab/toll-road/testutil/keeper"
	"github.com/b9lab/toll-road/x/tollroad"
	"github.com/b9lab/toll-road/x/tollroad/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRoadOperatorMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.TollroadKeeper(t)
	tollroad.InitGenesis(ctx, *k, *types.DefaultGenesis())
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 1; i < 6; i++ {
		expected := &types.MsgCreateRoadOperator{
			Creator: creator,
			Name:    strconv.Itoa(rand.Intn(1000)),
			Token:   strconv.Itoa(rand.Intn(1000)),
			Active:  true,
		}
		response, err := srv.CreateRoadOperator(wctx, expected)
		require.NoError(t, err)
		require.EqualValues(t, types.MsgCreateRoadOperatorResponse{
			Index: strconv.Itoa(i),
		}, *response)
		rst, found := k.GetRoadOperator(ctx,
			strconv.Itoa(i),
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
		sysInfo, found := k.GetSystemInfo(ctx)
		require.True(t, found)
		require.Equal(t, uint64(i+1), sysInfo.NextOperatorId)

		events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
		require.Len(t, events, 1)
		event := events[0]
		require.Equal(t, "new-road-operator-created", event.Type)
		attributes := event.Attributes
		require.EqualValues(t, []sdk.Attribute{
			{Key: "creator", Value: creator},
			{Key: "road-operator-index", Value: response.Index},
			{Key: "name", Value: expected.Name},
			{Key: "token", Value: expected.Token},
			{Key: "active", Value: "true"},
		}, attributes[(i-1)*5:i*5])

	}
}

func TestRoadOperatorMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateRoadOperator
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRoadOperator{Creator: creator,
				Index: strconv.Itoa(1),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRoadOperator{Creator: "B",
				Index: strconv.Itoa(1),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRoadOperator{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TollroadKeeper(t)
			tollroad.InitGenesis(ctx, *k, *types.DefaultGenesis())
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateRoadOperator{
				Creator: creator,
			}
			_, err := srv.CreateRoadOperator(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRoadOperator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRoadOperator(ctx,
					"1",
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRoadOperatorMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteRoadOperator
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRoadOperator{Creator: creator,
				Index: strconv.Itoa(1),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRoadOperator{Creator: "B",
				Index: strconv.Itoa(1),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRoadOperator{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.TollroadKeeper(t)
			tollroad.InitGenesis(ctx, *k, *types.DefaultGenesis())
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateRoadOperator(wctx, &types.MsgCreateRoadOperator{
				Creator: creator,
			})
			require.NoError(t, err)
			_, err = srv.DeleteRoadOperator(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRoadOperator(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
