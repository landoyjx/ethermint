package evm

import (
	"fmt"
	"math/big"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// BeginBlock sets the Bloom and Hash mappings and resets the Bloom filter and
// the transaction count to 0.
func BeginBlock(k Keeper, ctx sdk.Context, req abci.RequestBeginBlock) {
	if req.Header.LastBlockId.GetHash() == nil || req.Header.GetHeight() < 1 {
		return
	}

	// Consider removing this when using evm as module without web3 API
	bloom := ethtypes.BytesToBloom(k.Bloom.Bytes())
	k.SetBlockBloomMapping(ctx, bloom, req.Header.GetHeight()-1)
	k.SetBlockHashMapping(ctx, req.Header.LastBlockId.GetHash(), req.Header.GetHeight()-1)
	k.Bloom = big.NewInt(0)
	k.TxCount = 0
}

// EndBlock updates the accounts and commits states objects to the KV Store
func EndBlock(k Keeper, ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	snapshot := k.CommitStateDB.Snapshot()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error in ethermint end block", r);
			k.RevertToSnapshot(ctx, snapshot)
		}
	}()

	// Gas costs are handled within msg handler so costs should be ignored
	ctx = ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	currentGasMeter := ctx.GasMeter()

	fmt.Println("end block ,before update accounts")
	fmt.Println("gas meter10, ", "GasConsumed=", currentGasMeter.GasConsumed(),
		"GasConsumedToLimit=", currentGasMeter.GasConsumedToLimit(),
		"Limit=", currentGasMeter.Limit(),
		"IsPastLimit=", currentGasMeter.IsPastLimit(),
		"IsOutOfGas=", currentGasMeter.IsOutOfGas())
	// Update account balances before committing other parts of state
	k.CommitStateDB.UpdateAccounts()
	fmt.Println("end block ,after update accounts")
	fmt.Println("gas meter11, ", "GasConsumed=", currentGasMeter.GasConsumed(),
		"GasConsumedToLimit=", currentGasMeter.GasConsumedToLimit(),
		"Limit=", currentGasMeter.Limit(),
		"IsPastLimit=", currentGasMeter.IsPastLimit(),
		"IsOutOfGas=", currentGasMeter.IsOutOfGas())

	// Commit state objects to KV store
	_, err := k.CommitStateDB.WithContext(ctx).Commit(true)
	fmt.Println("commit state")
	fmt.Println("gas meter12, ", "GasConsumed=", currentGasMeter.GasConsumed(),
		"GasConsumedToLimit=", currentGasMeter.GasConsumedToLimit(),
		"Limit=", currentGasMeter.Limit(),
		"IsPastLimit=", currentGasMeter.IsPastLimit(),
		"IsOutOfGas=", currentGasMeter.IsOutOfGas())
	if err != nil {
		panic(err)
	}

	// Clear accounts cache after account data has been committed
	k.CommitStateDB.ClearStateObjects()

	return []abci.ValidatorUpdate{}
}
