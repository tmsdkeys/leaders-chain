package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tmsdkeys/leaders/x/leaders/types"
)

// SetTopRanked set topRanked in the store
func (k Keeper) SetTopRanked(ctx sdk.Context, topRanked types.TopRanked) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopRankedKey))
	b := k.cdc.MustMarshal(&topRanked)
	store.Set([]byte{0}, b)
}

// GetTopRanked returns topRanked
func (k Keeper) GetTopRanked(ctx sdk.Context) (val types.TopRanked, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopRankedKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTopRanked removes topRanked from the store
func (k Keeper) RemoveTopRanked(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TopRankedKey))
	store.Delete([]byte{0})
}
