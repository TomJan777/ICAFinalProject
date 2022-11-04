package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/b9lab/toll-road/testutil/keeper"
	"github.com/b9lab/toll-road/testutil/nullify"
	"github.com/b9lab/toll-road/x/tollroad/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRoadOperator(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.RoadOperator {
	items := make([]types.RoadOperator, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRoadOperator(ctx, items[i])
	}
	return items
}

func TestRoadOperatorGet(t *testing.T) {
	keeper, ctx := keepertest.TollroadKeeper(t)
	items := createNRoadOperator(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRoadOperator(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRoadOperatorRemove(t *testing.T) {
	keeper, ctx := keepertest.TollroadKeeper(t)
	items := createNRoadOperator(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRoadOperator(ctx,
			item.Index,
		)
		_, found := keeper.GetRoadOperator(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRoadOperatorGetAll(t *testing.T) {
	keeper, ctx := keepertest.TollroadKeeper(t)
	items := createNRoadOperator(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRoadOperator(ctx)),
	)
}
