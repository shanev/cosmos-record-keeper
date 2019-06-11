package recordkeeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RecordKeeper data type with a default codec
type RecordKeeper struct {
	storeKey sdk.StoreKey
	codec    *codec.Codec
}

// NewRecordKeeper creates a new record keeper for module keepers to embed
func NewRecordKeeper(storeKey sdk.StoreKey, codec *codec.Codec) RecordKeeper {
	return RecordKeeper{storeKey, codec}
}

// StoreKey returns the keeper store key.
func (k RecordKeeper) StoreKey() sdk.StoreKey {
	return k.storeKey
}
