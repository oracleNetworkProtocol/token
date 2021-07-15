package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupKeeper(t testing.TB) (*Keeper, sdk.Context) {
	// storeKey := sdk.NewKVStoreKey(types.StoreKey)
	// memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	// db := tmdb.NewMemDB()
	// stateStore := store.NewCommitMultiStore(db)
	// stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	// stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	// require.NoError(t, stateStore.LoadLatestVersion())
	// registry := codectypes.NewInterfaceRegistry()
	// keeper := NewKeeper(codec.NewProtoCodec(registry), storeKey, memStoreKey,)

	// ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	// return keeper, ctx
	return nil, sdk.Context{}
}
