package consumerchain2_test

import (
	"testing"

	keepertest "consumerchain2/testutil/keeper"
	"consumerchain2/testutil/nullify"
	"consumerchain2/x/consumerchain2"
	"consumerchain2/x/consumerchain2/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Consumerchain2Keeper(t)
	consumerchain2.InitGenesis(ctx, k, genesisState)
	got := consumerchain2.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
