package keeper

import (
	"github.com/b9lab/toll-road/x/tollroad/types"
)

var _ types.QueryServer = Keeper{}
