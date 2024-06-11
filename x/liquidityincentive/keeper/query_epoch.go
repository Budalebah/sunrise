package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sunriselayer/sunrise/x/liquidityincentive/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EpochAll(ctx context.Context, req *types.QueryAllEpochRequest) (*types.QueryAllEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var epochs []types.Epoch

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	epochStore := prefix.NewStore(store, types.KeyPrefix(types.EpochKey))

	pageRes, err := query.Paginate(epochStore, req.Pagination, func(key []byte, value []byte) error {
		var epoch types.Epoch
		if err := k.cdc.Unmarshal(value, &epoch); err != nil {
			return err
		}

		epochs = append(epochs, epoch)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEpochResponse{Epoch: epochs, Pagination: pageRes}, nil
}

func (k Keeper) Epoch(ctx context.Context, req *types.QueryGetEpochRequest) (*types.QueryGetEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	epoch, found := k.GetEpoch(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEpochResponse{Epoch: epoch}, nil
}
