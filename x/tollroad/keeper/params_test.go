package keeper_test

import (
	"testing"

	testkeeper "github.com/b9lab/toll-road/testutil/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TollroadKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
