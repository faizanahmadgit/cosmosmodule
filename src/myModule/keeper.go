package yourmodule

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define the keeper struct
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// Implement the NewKeeper function
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Implement the SetKeyValue function to store the key-value pair in the module's state
func (k Keeper) SetKeyValue(ctx sdk.Context, key, value string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(key), []byte(value))
}

// Implement the GetKeyValue function to retrieve the value of a given key from the module's state
func (k Keeper) GetKeyValue(ctx sdk.Context, key string) ([]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get([]byte(key))
	if value == nil {
		return nil, false
	}
	return value, true
}
