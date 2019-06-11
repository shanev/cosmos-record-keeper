package recordkeeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StringKeyedKeeper defines methods for the active record pattern
type StringKeyedKeeper interface {
	StringGet(ctx sdk.Context, key string, value interface{})
	StringSet(ctx sdk.Context, key string, value interface{})
}

// interface conformance check
var _ StringKeyedKeeper = RecordKeeper{}

// StringGet gets a value given a key
func (k RecordKeeper) StringGet(ctx sdk.Context, key string, value interface{}) {
	keyBytes := []byte(key)
	valueBytes := k.store(ctx).Get(keyBytes)
	if valueBytes == nil {
		return
	}
	k.codec.MustUnmarshalBinaryLengthPrefixed(valueBytes, value)
}

// StringSet sets a key, value pair in the store
func (k RecordKeeper) StringSet(ctx sdk.Context, key string, value interface{}) {
	valueBytes := k.codec.MustMarshalBinaryLengthPrefixed(value)
	k.stringSetBytes(ctx, key, valueBytes)
}

func (k RecordKeeper) stringSetBytes(ctx sdk.Context, key string, valueBytes []byte) {
	keyBytes := []byte(key)
	k.store(ctx).Set(keyBytes, valueBytes)
}
