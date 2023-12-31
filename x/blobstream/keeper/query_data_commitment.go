package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sunrise/x/blobstream/types"
)

func (k Keeper) DataCommitmentRangeForHeight(
	c context.Context,
	request *types.QueryDataCommitmentRangeForHeightRequest,
) (*types.QueryDataCommitmentRangeForHeightResponse, error) {
	resp, err := k.GetDataCommitmentForHeight(sdk.UnwrapSDKContext(c), request.Height)
	if err != nil {
		return nil, err
	}
	return &types.QueryDataCommitmentRangeForHeightResponse{DataCommitment: &resp}, nil
}

func (k Keeper) LatestDataCommitment(
	c context.Context,
	_ *types.QueryLatestDataCommitmentRequest,
) (*types.QueryLatestDataCommitmentResponse, error) {
	resp, err := k.GetLatestDataCommitment(sdk.UnwrapSDKContext(c))
	if err != nil {
		return nil, err
	}
	return &types.QueryLatestDataCommitmentResponse{
		DataCommitment: &resp,
	}, nil
}
