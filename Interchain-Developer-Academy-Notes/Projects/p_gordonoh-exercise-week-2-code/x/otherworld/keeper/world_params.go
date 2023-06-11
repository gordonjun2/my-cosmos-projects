package keeper

import (
	"github.com/b9lab/other-world/x/otherworld/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetWorldParams set worldParams in the store
func (k Keeper) SetWorldParams(ctx sdk.Context, worldParams types.WorldParams) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParamsKey))
	b := k.cdc.MustMarshal(&worldParams)
	store.Set([]byte{0}, b)
}

// GetWorldParams returns worldParams
func (k Keeper) GetWorldParams(ctx sdk.Context) (val types.WorldParams, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParamsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWorldParams removes worldParams from the store
func (k Keeper) RemoveWorldParams(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WorldParamsKey))
	store.Delete([]byte{0})
}
