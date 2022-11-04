package tollroad

import (
	"math/rand"

	"github.com/b9lab/toll-road/testutil/sample"
	tollroadsimulation "github.com/b9lab/toll-road/x/tollroad/simulation"
	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = tollroadsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateRoadOperator = "op_weight_msg_road_operator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRoadOperator int = 100

	opWeightMsgUpdateRoadOperator = "op_weight_msg_road_operator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRoadOperator int = 100

	opWeightMsgDeleteRoadOperator = "op_weight_msg_road_operator"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRoadOperator int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tollroadGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		RoadOperatorList: []types.RoadOperator{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tollroadGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRoadOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRoadOperator, &weightMsgCreateRoadOperator, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRoadOperator = defaultWeightMsgCreateRoadOperator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRoadOperator,
		tollroadsimulation.SimulateMsgCreateRoadOperator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRoadOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRoadOperator, &weightMsgUpdateRoadOperator, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRoadOperator = defaultWeightMsgUpdateRoadOperator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRoadOperator,
		tollroadsimulation.SimulateMsgUpdateRoadOperator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRoadOperator int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRoadOperator, &weightMsgDeleteRoadOperator, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRoadOperator = defaultWeightMsgDeleteRoadOperator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRoadOperator,
		tollroadsimulation.SimulateMsgDeleteRoadOperator(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
