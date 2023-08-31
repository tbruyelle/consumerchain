package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "consumerchain2/testutil/keeper"
	"consumerchain2/x/consumerchain2/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.Consumerchain2Keeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
