package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/token/x/token/keeper"
	"github.com/oracleNetworkProtocol/token/x/token/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the Token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, *elem)
	}

	// Set Token count
	k.SetTokenCount(ctx, genState.TokenCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all Token
	TokenList := k.GetAllToken(ctx)
	for _, elem := range TokenList {
		elem := elem
		genesis.TokenList = append(genesis.TokenList, &elem)
	}

	// Set the current count
	genesis.TokenCount = k.GetTokenCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
