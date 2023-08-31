package keeper

import (
	"consumerchain2/x/consumerchain2/types"
)

var _ types.QueryServer = Keeper{}
