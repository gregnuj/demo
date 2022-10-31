package keeper_test

import (
	"context"
	"testing"

	keepertest "demo/testutil/keeper"
	"demo/x/demo/keeper"
	"demo/x/demo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DemoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
