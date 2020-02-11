package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/internal/types"
	"github.com/stretchr/testify/require"
)

func cloneAppend(bz []byte, tail []byte) (res []byte) {
	res = make([]byte, len(bz)+len(tail))
	copy(res, bz)
	copy(res[len(bz):], tail)
	return
}

func TestAddressFromBalancesStore(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("cosmos1n88uc38xhjgxzw9nwre4ep2c8ga4fjxcar6mn7")
	require.NoError(t, err)

	key := cloneAppend(addr.Bytes(), []byte("stake"))
	res := types.AddressFromBalancesStore(key)
	require.Equal(t, res, addr)
}
