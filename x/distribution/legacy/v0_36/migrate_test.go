package v0_36

import (
	"github.com/osiz-blockchainapp/bitcoiva-sdk/types"
	v034distr "github.com/osiz-blockchainapp/bitcoiva-sdk/x/distribution/legacy/v0_34"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	priv       = secp256k1.GenPrivKey()
	addr       = types.AccAddress(priv.PubKey().Address())
	valAddr, _ = types.ValAddressFromBech32(addr.String())

	event = v034distr.ValidatorSlashEvent{
		ValidatorPeriod: 1,
		Fraction:        types.Dec{},
	}
)

func TestMigrate(t *testing.T) {
	var genesisState GenesisState
	require.NotPanics(t, func() {
		genesisState = Migrate(v034distr.GenesisState{
			ValidatorSlashEvents: []v034distr.ValidatorSlashEventRecord{
				{
					ValidatorAddress: valAddr,
					Height:           1,
					Event:            event,
				},
			},
		})
	})

	require.Equal(t, genesisState.ValidatorSlashEvents[0], ValidatorSlashEventRecord{
		ValidatorAddress: valAddr,
		Height:           1,
		Period:           event.ValidatorPeriod,
		Event:            event,
	})
}

func TestMigrateEmptyRecord(t *testing.T) {
	var genesisState GenesisState

	require.NotPanics(t, func() {
		genesisState = Migrate(v034distr.GenesisState{
			ValidatorSlashEvents: []v034distr.ValidatorSlashEventRecord{{}},
		})
	})

	require.Equal(t, genesisState.ValidatorSlashEvents[0], ValidatorSlashEventRecord{
		ValidatorAddress: valAddr,
		Height:           0,
		Period:           0,
		Event: v034distr.ValidatorSlashEvent{
			ValidatorPeriod: 0,
			Fraction:        types.Dec{},
		},
	})
}
