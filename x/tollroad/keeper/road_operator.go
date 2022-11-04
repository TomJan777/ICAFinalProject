package keeper

import (
	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRoadOperator set a specific roadOperator in the store from its index
func (k Keeper) SetRoadOperator(ctx sdk.Context, roadOperator types.RoadOperator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoadOperatorKeyPrefix))
	b := k.cdc.MustMarshal(&roadOperator)
	store.Set(types.RoadOperatorKey(
		roadOperator.Index,
	), b)
}

// GetRoadOperator returns a roadOperator from its index
func (k Keeper) GetRoadOperator(
	ctx sdk.Context,
	index string,

) (val types.RoadOperator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoadOperatorKeyPrefix))

	b := store.Get(types.RoadOperatorKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRoadOperator removes a roadOperator from the store
func (k Keeper) RemoveRoadOperator(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoadOperatorKeyPrefix))
	store.Delete(types.RoadOperatorKey(
		index,
	))
}

// GetAllRoadOperator returns all roadOperator
func (k Keeper) GetAllRoadOperator(ctx sdk.Context) (list []types.RoadOperator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RoadOperatorKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.RoadOperator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
