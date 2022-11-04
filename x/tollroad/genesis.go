package tollroad

import (
	"github.com/b9lab/toll-road/x/tollroad/keeper"
	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetSystemInfo(ctx, genState.SystemInfo)
	// Set all the roadOperator
	for _, elem := range genState.RoadOperatorList {
		k.SetRoadOperator(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = systemInfo
	}
	genesis.RoadOperatorList = k.GetAllRoadOperator(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
