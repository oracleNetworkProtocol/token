package keeper

import (
	"context"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/oracleNetworkProtocol/token/x/token/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var Token = types.Token{
		Creator:        msg.Creator,
		Symbol:         msg.Symbol,
		OriginalSymbol: msg.OriginalSymbol,
		Desc:           msg.Desc,
		WholeName:      msg.WholeName,
		TotalSupply:    msg.TotalSupply,
		Own:            msg.Own,
		Mintable:       msg.Mintable,
	}

	id := k.AppendToken(
		ctx,
		Token,
	)
	_, err := k.MintAndCreateToken(ctx, Token)
	if err != nil {
		log.Println("mint and create token err:", err)
		return &types.MsgCreateTokenResponse{}, nil
	}
	k.Keeper.SetToken(ctx, Token)

	return &types.MsgCreateTokenResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var Token = types.Token{
		Creator:        msg.Creator,
		Id:             msg.Id,
		Symbol:         msg.Symbol,
		OriginalSymbol: msg.OriginalSymbol,
		Desc:           msg.Desc,
		WholeName:      msg.WholeName,
		TotalSupply:    msg.TotalSupply,
		Own:            msg.Own,
		Mintable:       msg.Mintable,
	}

	// Checks that the element exists
	if !k.HasToken(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetTokenOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetToken(ctx, Token)

	return &types.MsgUpdateTokenResponse{}, nil
}

func (k msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasToken(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetTokenOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveToken(ctx, msg.Id)

	return &types.MsgDeleteTokenResponse{}, nil
}
