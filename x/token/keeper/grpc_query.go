package keeper

import (
	"github.com/oracleNetworkProtocol/token/x/token/types"
)

var _ types.QueryServer = Keeper{}
