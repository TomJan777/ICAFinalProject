package mock_types

import (
	"context"

	"github.com/b9lab/toll-road/x/tollroad/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (escrow *MockBankEscrowKeeper) ExpectAny(context context.Context) {
	escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
}

func coinsOf(token string, amount uint64) sdk.Coins {
	return sdk.Coins{
		sdk.Coin{
			Denom:  token,
			Amount: sdk.NewInt(int64(amount)),
		},
	}
}

func (escrow *MockBankEscrowKeeper) ExpectFundVault(context context.Context, who string, token string, amount uint64) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), whoAddr, types.ModuleName, coinsOf(token, amount))
}

func (escrow *MockBankEscrowKeeper) ExpectRefundVault(context context.Context, who string, token string, amount uint64) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), types.ModuleName, whoAddr, coinsOf(token, amount))
}
