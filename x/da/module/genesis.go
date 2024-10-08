package da

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sunriselayer/sunrise/x/da/keeper"
	"github.com/sunriselayer/sunrise/x/da/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
	for _, data := range genState.PublishedData {
		if err := k.SetPublishedData(ctx, data); err != nil {
			panic(err)
		}
	}
	for _, proof := range genState.Proofs {
		if err := k.SetProof(ctx, proof); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.PublishedData = k.GetAllPublishedData(ctx)
	genesis.Proofs = k.GetAllProofs(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
