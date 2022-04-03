package cmd_test

import (
	"context"
	"fmt"
	"testing"

	matrixd "github.com/MatrixDao/matrix/cmd/matrixd/cmd"
	"github.com/MatrixDao/matrix/x/testutil/sample"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

var testModuleBasicManager = module.NewBasicManager(genutil.AppModuleBasic{})

// Tests "add-genesis-account", a command that adds a genesis account to genesis.json
func TestAddGenesisAccountCmd(t *testing.T) {

	type TestCase struct {
		name        string
		addr        string
		denom       string
		expectError bool
	}

	var executeTest = func(t *testing.T, testCase TestCase) {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {

			home := t.TempDir()
			logger := log.NewNopLogger()
			cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
			require.NoError(t, err)

			appCodec := simapp.MakeTestEncodingConfig().Marshaler
			err = genutiltest.ExecInitCmd(
				testModuleBasicManager, home, appCodec)
			require.NoError(t, err)

			serverCtx := server.NewContext(viper.New(), cfg, logger)
			clientCtx := client.Context{}.WithCodec(appCodec).WithHomeDir(home)

			ctx := context.Background()
			ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
			ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)

			cmd := matrixd.AddGenesisAccountCmd(home)
			cmd.SetArgs([]string{
				tc.addr,
				tc.denom,
				fmt.Sprintf("--%s=home", flags.FlagHome)})

			if tc.expectError {
				require.Error(t, cmd.ExecuteContext(ctx))
			} else {
				require.NoError(t, cmd.ExecuteContext(ctx))
			}
		})
	}

	sampleAddr := sample.AccAddress()

	testCases := []TestCase{
		{
			name:        "invalid address",
			addr:        "",
			denom:       "1000atom",
			expectError: true,
		},
		{
			name:        "valid address",
			addr:        sampleAddr.String(),
			denom:       "1000atom",
			expectError: false,
		},
		{
			name:        "multiple denoms",
			addr:        sampleAddr.String(),
			denom:       "1000atom, 2000stake",
			expectError: false,
		},
	}

	for _, testCase := range testCases {
		executeTest(t, testCase)
	}
}