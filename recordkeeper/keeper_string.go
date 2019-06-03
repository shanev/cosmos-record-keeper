package recordkeeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StringIterableKeeper defines methods for the active record pattern
type StringIterableKeeper interface {
	Get(ctx sdk.Context, key string, value interface{})
	Set(ctx sdk.Context, key string, value interface{})
}

// interface conformance check
var _ StringIterableKeeper = StringRecordKeeper{}

// StringRecordKeeper data type with a default codec
type StringRecordKeeper struct {
	storeKey sdk.StoreKey
	codec    *codec.Codec
}

// NewStringRecordKeeper creates a new record keeper for module keepers to embed
func NewStringRecordKeeper(storeKey sdk.StoreKey, codec *codec.Codec) StringRecordKeeper {
	return StringRecordKeeper{storeKey, codec}
}

// Get gets a value given a key
func (k StringRecordKeeper) Get(ctx sdk.Context, key string, value interface{}) {
	keyBytes := []byte(key)
	valueBytes := k.store(ctx).Get(keyBytes)
	if valueBytes == nil {
		return
	}
	k.codec.MustUnmarshalBinaryLengthPrefixed(valueBytes, value)
}

// Set sets a key, value pair in the store
func (k StringRecordKeeper) Set(ctx sdk.Context, key string, value interface{}) {
	valueBytes := k.codec.MustMarshalBinaryLengthPrefixed(value)
	k.setBytes(ctx, key, valueBytes)
}

// Internal

func (k StringRecordKeeper) setBytes(ctx sdk.Context, key string, valueBytes []byte) {
	keyBytes := []byte(key)
	k.store(ctx).Set(keyBytes, valueBytes)
}

func (k StringRecordKeeper) store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}
