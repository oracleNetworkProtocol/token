package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/oracleNetworkProtocol/token/x/token/types"
	"strconv"
)

// GetTokenCount get the total number of Token
func (k Keeper) GetTokenCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenCountKey))
	byteKey := types.KeyPrefix(types.TokenCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetTokenCount set the total number of Token
func (k Keeper) SetTokenCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenCountKey))
	byteKey := types.KeyPrefix(types.TokenCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendToken appends a Token in the store with a new id and update the count
func (k Keeper) AppendToken(
	ctx sdk.Context,
	Token types.Token,
) uint64 {
	// Create the Token
	count := k.GetTokenCount(ctx)

	// Set the ID of the appended value
	Token.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&Token)
	store.Set(GetTokenIDBytes(Token.Id), appendedValue)

	// Update Token count
	k.SetTokenCount(ctx, count+1)

	return count
}

// SetToken set a specific Token in the store
func (k Keeper) SetToken(ctx sdk.Context, Token types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	b := k.cdc.MustMarshalBinaryBare(&Token)
	store.Set(GetTokenIDBytes(Token.Id), b)
}

// GetToken returns a Token from its id
func (k Keeper) GetToken(ctx sdk.Context, id uint64) types.Token {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	var Token types.Token
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetTokenIDBytes(id)), &Token)
	return Token
}

// HasToken checks if the Token exists in the store
func (k Keeper) HasToken(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	return store.Has(GetTokenIDBytes(id))
}

// GetTokenOwner returns the creator of the Token
func (k Keeper) GetTokenOwner(ctx sdk.Context, id uint64) string {
	return k.GetToken(ctx, id).Creator
}

// RemoveToken removes a Token from the store
func (k Keeper) RemoveToken(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	store.Delete(GetTokenIDBytes(id))
}

// GetAllToken returns all Token
func (k Keeper) GetAllToken(ctx sdk.Context) (list []types.Token) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Token
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTokenIDBytes returns the byte representation of the ID
func GetTokenIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTokenIDFromBytes returns ID in uint64 format from a byte array
func GetTokenIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
