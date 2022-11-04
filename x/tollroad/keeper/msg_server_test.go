package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/b9lab/toll-road/testutil/keeper"
	"github.com/b9lab/toll-road/x/tollroad/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TollroadKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
