package types_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	appparams "github.com/mokitanetwork/katana/app/params"
	"github.com/mokitanetwork/katana/x/leverage/types"
)

func TestAddressFromKey(t *testing.T) {
	address := sdk.AccAddress([]byte("addr________________"))
	key := types.KeyAdjustedBorrow(address, appparams.BondDenom)
	expectedAddress := types.AddressFromKey(key, types.KeyPrefixAdjustedBorrow)

	require.Equal(t, address, expectedAddress)

	address = sdk.AccAddress([]byte("anotherAddr________________"))
	key = types.KeyCollateralAmountNoDenom(address)
	expectedAddress = types.AddressFromKey(key, types.KeyPrefixAdjustedBorrow)

	require.Equal(t, address, expectedAddress)
}

func TestDenomFromKeyWithAddress(t *testing.T) {
	address := sdk.AccAddress([]byte("addr________________"))
	denom := appparams.BondDenom
	key := types.KeyAdjustedBorrow(address, denom)
	expectedDenom := types.DenomFromKeyWithAddress(key, types.KeyPrefixAdjustedBorrow)

	require.Equal(t, denom, expectedDenom)

	uDenom := fmt.Sprintf("u%s", denom)
	key = types.KeyCollateralAmount(address, uDenom)
	expectedDenom = types.DenomFromKeyWithAddress(key, types.KeyPrefixCollateralAmount)

	require.Equal(t, uDenom, expectedDenom)
}

func TestDenomFromKey(t *testing.T) {
	denom := appparams.BondDenom
	key := types.KeyReserveAmount(denom)
	expectedDenom := types.DenomFromKey(key, types.KeyPrefixReserveAmount)

	require.Equal(t, denom, expectedDenom)

	uDenom := fmt.Sprintf("u%s", denom)
	key = types.KeyReserveAmount(uDenom)
	expectedDenom = types.DenomFromKey(key, types.KeyPrefixReserveAmount)

	require.Equal(t, uDenom, expectedDenom)
}

func TestGetKeys(t *testing.T) {
	type testCase struct {
		actual      []byte
		expected    [][]byte
		description string
	}

	addr := sdk.AccAddress("addr________________") // length: 20
	addrbytes := []byte{0x61, 0x64, 0x64, 0x72, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f, 0x5f}
	ukatanabytes := []byte{0x75, 0x75, 0x6d, 0x65, 0x65}                              // ukatana
	ibcabcdbytes := []byte{0x69, 0x62, 0x63, 0x2f, 0x61, 0x62, 0x63, 0x64}          // ibc/abcd
	uibcbytes := []byte{0x75, 0x2f, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x62, 0x63, 0x64} // u/ibc/abcd

	testCases := []testCase{
		{
			types.KeyRegisteredToken("ukatana"),
			[][]byte{
				{0x01},     // prefix
				ukatanabytes, // ukatana
				{0x00},     // null terminator
			},
			"registered token key (ukatana)",
		},
		{
			types.KeyRegisteredToken("ibc/abcd"),
			[][]byte{
				{0x01},       // prefix
				ibcabcdbytes, // ibc/abcd
				{0x00},       // null terminator
			},
			"registered token key (ibc/abcd)",
		},
		{
			types.KeyAdjustedBorrow(addr, "ukatana"),
			[][]byte{
				{0x02},     // prefix
				{0x14},     // address length prefix = 20
				addrbytes,  // addr________________
				ukatanabytes, // ukatana
				{0x00},     // null terminator
			},
			"adjusted borrow key (ukatana)",
		},
		{
			types.KeyAdjustedBorrow(addr, "ibc/abcd"),
			[][]byte{
				{0x02},       // prefix
				{0x14},       // address length prefix = 20
				addrbytes,    // addr________________
				ibcabcdbytes, // ibc/abcd
				{0x00},       // null terminator
			},
			"adjusted borrow key (ibc)",
		},
		{
			types.KeyAdjustedBorrowNoDenom(addr),
			[][]byte{
				{0x02},    // prefix
				{0x14},    // address length prefix = 20
				addrbytes, // addr________________
			},
			"adjusted borrow key (no denom)",
		},
		{
			types.KeyCollateralAmount(addr, "u/ibc/abcd"),
			[][]byte{
				{0x04},    // prefix
				{0x14},    // address length prefix = 20
				addrbytes, // addr________________
				uibcbytes, // u/ibc/abcd
				{0x00},    // null terminator
			},
			"collateral amount key",
		},
		{
			types.KeyCollateralAmountNoDenom(addr),
			[][]byte{
				{0x04},    // prefix
				{0x14},    // address length prefix = 20
				addrbytes, // addr________________
			},
			"collateral amount key (no denom)",
		},
		{
			types.KeyReserveAmount("ibc/abcd"),
			[][]byte{
				{0x05},       // prefix
				ibcabcdbytes, // ibc/abcd
				{0x00},       // null terminator
			},
			"reserve amount key",
		},
		{
			types.KeyBadDebt("u/ibc/abcd", addr),
			[][]byte{
				{0x07},    // prefix
				{0x14},    // address length prefix = 20
				addrbytes, // addr________________
				uibcbytes, // u/ibc/abcd
				{0x00},    // null terminator
			},
			"bad debt key",
		},
		{
			types.KeyInterestScalar("ibc/abcd"),
			[][]byte{
				{0x08},       // prefix
				ibcabcdbytes, // ibc/abcd
				{0x00},       // null terminator
			},
			"interest scalar key",
		},
		{
			types.KeyAdjustedTotalBorrow("ibc/abcd"),
			[][]byte{
				{0x09},       // prefix
				ibcabcdbytes, // ibc/abcd
				{0x00},       // null terminator
			},
			"adjusted total borrow key",
		},
		{
			types.KeyUTokenSupply("u/ibc/abcd"),
			[][]byte{
				{0x0A},    // prefix
				uibcbytes, // u/ibc/abcd
				{0x00},    // null terminator
			},
			"uToken supply key",
		},
	}
	for _, tc := range testCases {
		expectedKey := []byte{}
		for _, e := range tc.expected {
			expectedKey = append(expectedKey, e...)
		}
		require.Equalf(
			t,
			expectedKey,
			tc.actual,
			tc.description,
		)
	}
}
