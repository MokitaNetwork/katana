package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mokitanetwork/katana/x/leverage/types"
)

func (s *IntegrationTestSuite) TestGetBorrow() {
	ctx, require := s.ctx, s.Require()

	// get ukatana borrow amount of empty account address (zero)
	borrowed := s.tk.GetBorrow(ctx, sdk.AccAddress{}, katanaDenom)
	require.Equal(coin(katanaDenom, 0), borrowed)

	// creates account which has supplied and collateralized 1000 ukatana, and borrowed 123 ukatana
	addr := s.newAccount(coin(katanaDenom, 1000))
	s.supply(addr, coin(katanaDenom, 1000))
	s.collateralize(addr, coin("u/"+katanaDenom, 1000))
	s.borrow(addr, coin(katanaDenom, 123))

	// confirm borrowed amount is 123 ukatana
	borrowed = s.tk.GetBorrow(ctx, addr, katanaDenom)
	require.Equal(coin(katanaDenom, 123), borrowed)

	// unregistered denom (zero)
	borrowed = s.tk.GetBorrow(ctx, addr, "abcd")
	require.Equal(coin("abcd", 0), borrowed)

	// we do not test empty denom, as that will cause a panic
}

func (s *IntegrationTestSuite) TestSetBorrow() {
	ctx, require := s.ctx, s.Require()

	// empty account address
	err := s.tk.SetBorrow(ctx, sdk.AccAddress{}, coin(katanaDenom, 123))
	require.ErrorIs(err, types.ErrEmptyAddress)

	addr := s.newAccount()

	// set nonzero borrow, and confirm effect
	err = s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 123))
	require.NoError(err)
	require.Equal(coin(katanaDenom, 123), s.tk.GetBorrow(ctx, addr, katanaDenom))

	// set zero borrow, and confirm effect
	err = s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 0))
	require.NoError(err)
	require.Equal(coin(katanaDenom, 0), s.tk.GetBorrow(ctx, addr, katanaDenom))

	// unregistered (but valid) denom
	err = s.tk.SetBorrow(ctx, addr, coin("abcd", 123))
	require.NoError(err)

	// interest scalar test - ensure borrowing smallest possible amount doesn't round to zero at scalar = 1.0001
	require.NoError(s.tk.SetInterestScalar(ctx, katanaDenom, sdk.MustNewDecFromStr("1.0001")))
	require.NoError(s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 1)))
	require.Equal(coin(katanaDenom, 1), s.tk.GetBorrow(ctx, addr, katanaDenom))

	// interest scalar test - scalar changing after borrow (as it does when interest accrues)
	require.NoError(s.tk.SetInterestScalar(ctx, katanaDenom, sdk.MustNewDecFromStr("1.0")))
	require.NoError(s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 200)))
	require.NoError(s.tk.SetInterestScalar(ctx, katanaDenom, sdk.MustNewDecFromStr("2.33")))
	require.Equal(coin(katanaDenom, 466), s.tk.GetBorrow(ctx, addr, katanaDenom))

	// interest scalar extreme case - rounding up becomes apparent at high borrow amount
	require.NoError(s.tk.SetInterestScalar(ctx, katanaDenom, sdk.MustNewDecFromStr("555444333222111")))
	require.NoError(s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 1)))
	require.Equal(coin(katanaDenom, 1), s.tk.GetBorrow(ctx, addr, katanaDenom))
	require.NoError(s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 20000)))
	require.Equal(coin(katanaDenom, 20001), s.tk.GetBorrow(ctx, addr, katanaDenom))
}

func (s *IntegrationTestSuite) TestGetTotalBorrowed() {
	ctx, require := s.ctx, s.Require()

	// unregistered denom (zero)
	borrowed := s.tk.GetTotalBorrowed(ctx, "abcd")
	require.Equal(coin("abcd", 0), borrowed)

	// creates account which has supplied and collateralized 1000 ukatana, and borrowed 123 ukatana
	borrower := s.newAccount(coin(katanaDenom, 1000))
	s.supply(borrower, coin(katanaDenom, 1000))
	s.collateralize(borrower, coin("u/"+katanaDenom, 1000))
	s.borrow(borrower, coin(katanaDenom, 123))

	// confirm total borrowed amount is 123 ukatana
	borrowed = s.tk.GetTotalBorrowed(ctx, katanaDenom)
	require.Equal(coin(katanaDenom, 123), borrowed)

	// creates account which has supplied and collateralized 1000 ukatana, and borrowed 234 ukatana
	borrower2 := s.newAccount(coin(katanaDenom, 1000))
	s.supply(borrower2, coin(katanaDenom, 1000))
	s.collateralize(borrower2, coin("u/"+katanaDenom, 1000))
	s.borrow(borrower2, coin(katanaDenom, 234))

	// confirm total borrowed amount is 357 ukatana
	borrowed = s.tk.GetTotalBorrowed(ctx, katanaDenom)
	require.Equal(coin(katanaDenom, 357), borrowed)

	// interest scalar test - scalar changing after borrow (as it does when interest accrues)
	require.NoError(s.tk.SetInterestScalar(ctx, katanaDenom, sdk.MustNewDecFromStr("2.00")))
	require.Equal(coin(katanaDenom, 714), s.tk.GetTotalBorrowed(ctx, katanaDenom))
}

func (s *IntegrationTestSuite) TestLiquidity() {
	ctx, require := s.ctx, s.Require()

	// unregistered denom (zero)
	available := s.tk.AvailableLiquidity(ctx, "abcd")
	require.Equal(sdk.ZeroInt(), available)

	// creates account which has supplied and collateralized 1000 ukatana
	supplier := s.newAccount(coin(katanaDenom, 1000))
	s.supply(supplier, coin(katanaDenom, 1000))
	s.collateralize(supplier, coin("u/"+katanaDenom, 1000))

	// confirm lending pool is 1000 ukatana
	available = s.tk.AvailableLiquidity(ctx, katanaDenom)
	require.Equal(sdk.NewInt(1000), available)

	// creates account which has supplied and collateralized 1000 ukatana, and borrowed 123 ukatana
	borrower := s.newAccount(coin(katanaDenom, 1000))
	s.supply(borrower, coin(katanaDenom, 1000))
	s.collateralize(borrower, coin("u/"+katanaDenom, 1000))
	s.borrow(borrower, coin(katanaDenom, 123))

	// confirm lending pool is 1877 ukatana
	available = s.tk.AvailableLiquidity(ctx, katanaDenom)
	require.Equal(sdk.NewInt(1877), available)

	// artificially reserve 200 ukatana, reducing available amount to 1677
	s.setReserves(coin(katanaDenom, 200))
	available = s.tk.AvailableLiquidity(ctx, katanaDenom)
	require.Equal(sdk.NewInt(1677), available)
}

func (s *IntegrationTestSuite) TestDeriveBorrowUtilization() {
	ctx, require := s.ctx, s.Require()

	// unregistered denom (0 borrowed and 0 lending pool is considered 100%)
	utilization := s.tk.SupplyUtilization(ctx, "abcd")
	require.Equal(sdk.OneDec(), utilization)

	// creates account which has supplied and collateralized 1000 ukatana
	addr := s.newAccount(coin(katanaDenom, 1000))
	s.supply(addr, coin(katanaDenom, 1000))
	s.collateralize(addr, coin("u/"+katanaDenom, 1000))

	// All tests below are commented with the following equation in mind:
	//   utilization = (Total Borrowed / (Total Borrowed + Module Balance - Reserved Amount))

	// 0% utilization (0 / 0+1000-0)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.ZeroDec(), utilization)

	// user borrows 200 ukatana, reducing module account to 800 ukatana
	s.borrow(addr, coin(katanaDenom, 200))

	// 20% utilization (200 / 200+800-0)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("0.2"), utilization)

	// artificially reserve 200 ukatana
	s.setReserves(coin(katanaDenom, 200))

	// 25% utilization (200 / 200+800-200)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("0.25"), utilization)

	// user borrows 600 ukatana (disregard borrow limit), reducing module account to 0 ukatana
	s.forceBorrow(addr, coin(katanaDenom, 600))

	// 100% utilization (800 / 800+200-200))
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("1.0"), utilization)

	// artificially set user borrow to 1200 katana
	require.NoError(s.tk.SetBorrow(ctx, addr, coin(katanaDenom, 1200)))

	// still 100% utilization (1200 / 1200+200-200)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("1.0"), utilization)

	// artificially set reserves to 800 ukatana
	s.setReserves(coin(katanaDenom, 800))

	// edge case interpreted as 100% utilization (1200 / 1200+200-800)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("1.0"), utilization)

	// artificially set reserves to 4000 ukatana
	s.setReserves(coin(katanaDenom, 4000))

	// impossible case interpreted as 100% utilization (1200 / 1200+200-4000)
	utilization = s.tk.SupplyUtilization(ctx, katanaDenom)
	require.Equal(sdk.MustNewDecFromStr("1.0"), utilization)
}

func (s *IntegrationTestSuite) TestCalculateBorrowLimit() {
	app, ctx, require := s.app, s.ctx, s.Require()

	// Empty coins
	borrowLimit, err := app.LeverageKeeper.CalculateBorrowLimit(ctx, sdk.NewCoins())
	require.NoError(err)
	require.Equal(sdk.ZeroDec(), borrowLimit)

	// Unregistered asset
	invalidCoins := sdk.NewCoins(coin("abcd", 1000))
	_, err = app.LeverageKeeper.CalculateBorrowLimit(ctx, invalidCoins)
	require.ErrorIs(err, types.ErrNotUToken)

	// Create collateral uTokens (1k u/katana)
	katanaCollatDenom := types.ToUTokenDenom(katanaDenom)
	katanaCollateral := sdk.NewCoins(coin(katanaCollatDenom, 1000_000000))

	// Manually compute borrow limit using collateral weight of 0.25
	// and placeholder of 1 katana = $4.21.
	expectedKatanaLimit := sdk.NewDecFromInt(katanaCollateral[0].Amount).
		Mul(sdk.MustNewDecFromStr("0.00000421")).
		Mul(sdk.MustNewDecFromStr("0.25"))

	// Check borrow limit vs. manually computed value
	borrowLimit, err = app.LeverageKeeper.CalculateBorrowLimit(ctx, katanaCollateral)
	require.NoError(err)
	require.Equal(expectedKatanaLimit, borrowLimit)

	// Create collateral atom uTokens (1k u/uatom)
	atomCollatDenom := types.ToUTokenDenom(atomDenom)
	atomCollateral := sdk.NewCoins(coin(atomCollatDenom, 1000_000000))

	// Manually compute borrow limit using collateral weight of 0.25
	// and placeholder of 1 atom = $39.38
	expectedAtomLimit := sdk.NewDecFromInt(atomCollateral[0].Amount).
		Mul(sdk.MustNewDecFromStr("0.00003938")).
		Mul(sdk.MustNewDecFromStr("0.25"))

	// Check borrow limit vs. manually computed value
	borrowLimit, err = app.LeverageKeeper.CalculateBorrowLimit(ctx, atomCollateral)
	require.NoError(err)
	require.Equal(expectedAtomLimit, borrowLimit)

	// Compute the expected borrow limit of the two combined collateral coins
	expectedCombinedLimit := expectedKatanaLimit.Add(expectedAtomLimit)
	combinedCollateral := katanaCollateral.Add(atomCollateral...)

	// Check borrow limit vs. manually computed value
	borrowLimit, err = app.LeverageKeeper.CalculateBorrowLimit(ctx, combinedCollateral)
	require.NoError(err)
	require.Equal(expectedCombinedLimit, borrowLimit)
}
