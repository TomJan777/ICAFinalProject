package keeper

import (
	"context"

	"github.com/b9lab/toll-road/x/tollroad/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RoadOperatorAll(c context.Context, req *types.QueryAllRoadOperatorRequest) (*types.QueryAllRoadOperatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var roadOperators []types.RoadOperator
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	roadOperatorStore := prefix.NewStore(store, types.KeyPrefix(types.RoadOperatorKeyPrefix))

	pageRes, err := query.Paginate(roadOperatorStore, req.Pagination, func(key []byte, value []byte) error {
		var roadOperator types.RoadOperator
		if err := k.cdc.Unmarshal(value, &roadOperator); err != nil {
			return err
		}

		roadOperators = append(roadOperators, roadOperator)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRoadOperatorResponse{RoadOperator: roadOperators, Pagination: pageRes}, nil
}

func (k Keeper) RoadOperator(c context.Context, req *types.QueryGetRoadOperatorRequest) (*types.QueryGetRoadOperatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRoadOperator(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRoadOperatorResponse{RoadOperator: val}, nil
}
