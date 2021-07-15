package keeper

import (
	"fmt"
	golog "log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/oracleNetworkProtocol/token/x/token/types"
	"github.com/tendermint/tendermint/libs/log"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		paramSpace    paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	ak types.AccountKeeper, bk types.BankKeeper, ps paramtypes.Subspace,
) *Keeper {
	// ensure liquidity module account is set
	addr := ak.GetModuleAddress(types.ModuleName)
	golog.Println("module account address:", addr.String())
	if addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		accountKeeper: ak,
		bankKeeper:    bk,
		paramSpace:    ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintAndCreateToken(ctx sdk.Context, token types.Token) (sdk.Coin, error) {
	golog.Println("mint-and-create-token: 进入")
	total := sdk.NewInt(token.TotalSupply)
	golog.Println("total: ", total)

	mintCoin := sdk.NewCoin(token.Symbol, total)
	golog.Println("mintCoin: ", mintCoin)

	mintCoins := sdk.NewCoins(mintCoin)
	golog.Println("mintCoins: ", mintCoins)

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		golog.Println("k.Bank", err)
		return sdk.Coin{}, err
	}
	golog.Println("creator address:", token.Creator)
	adkAcc, er := sdk.AccAddressFromBech32(token.Creator)
	golog.Println("create address bech32:", adkAcc, er)
	if er != nil {
		return sdk.Coin{}, er
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, adkAcc, mintCoins); err != nil {
		golog.Println("k.sendCoinsFromModuleTo", err)
		return sdk.Coin{}, err
	}
	return mintCoin, nil
}
