package leaders

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tmsdkeys/leaders/x/leaders/keeper"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.TopRanked != nil {
		k.SetTopRanked(ctx, *genState.TopRanked)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	// Get all topRanked
	topRanked, found := k.GetTopRanked(ctx)
	if found {
		genesis.TopRanked = &topRanked
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
