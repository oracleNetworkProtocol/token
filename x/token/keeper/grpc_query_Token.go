package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/oracleNetworkProtocol/token/x/token/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenAll(c context.Context, req *types.QueryAllTokenRequest) (*types.QueryAllTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Tokens []*types.Token
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	TokenStore := prefix.NewStore(store, types.KeyPrefix(types.TokenKey))

	pageRes, err := query.Paginate(TokenStore, req.Pagination, func(key []byte, value []byte) error {
		var Token types.Token
		if err := k.cdc.UnmarshalBinaryBare(value, &Token); err != nil {
			return err
		}

		Tokens = append(Tokens, &Token)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTokenResponse{Token: Tokens, Pagination: pageRes}, nil
}

func (k Keeper) Token(c context.Context, req *types.QueryGetTokenRequest) (*types.QueryGetTokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Token types.Token
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasToken(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetTokenIDBytes(req.Id)), &Token)

	return &types.QueryGetTokenResponse{Token: &Token}, nil
}
