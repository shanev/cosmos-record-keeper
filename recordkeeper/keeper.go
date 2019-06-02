package recordkeeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RecordKeeper data type with a default codec
type RecordKeeper struct {
	storeKey sdk.StoreKey
}

// NewRecordKeeper creates a new record keeper for module keepers to embed
func NewRecordKeeper(storeKey sdk.StoreKey) RecordKeeper {
	return RecordKeeper{storeKey}
}

// StoreKey returns the default store key for the keeper
func (k RecordKeeper) StoreKey() sdk.StoreKey {
	return k.storeKey
}

// Store returns the default KVStore for the keeper
func (k RecordKeeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.StoreKey())
}

// NextID increments and returns the next available id by 1
func (k RecordKeeper) NextID(ctx sdk.Context) (id uint64) {
	idBytes := k.Store(ctx).Get(k.lenKey())
	if idBytes == nil {
		initialIndex := uint64(1)
		k.SetLen(ctx, initialIndex)

		return initialIndex
	}

	err := json.Unmarshal(idBytes, &id)
	if err != nil {
		panic(err)
	}

	nextID := id + 1
	k.SetLen(ctx, nextID)

	return nextID
}

// SetLen sets the len metadata in the store for incrementing ids
func (k RecordKeeper) SetLen(ctx sdk.Context, len uint64) {
	idBytes, err := json.Marshal(len)
	if err != nil {
		panic(err)
	}
	k.Store(ctx).Set(k.lenKey(), idBytes)
}

func (k RecordKeeper) lenKey() []byte {
	return []byte(k.storeKey.Name() + ":len")
}

// EachPrefix calls `fn` for each record in a store with a given prefix. Iteration will stop if `fn` returns false
func (k RecordKeeper) EachPrefix(ctx sdk.Context, prefix string, fn func([]byte) bool) (err sdk.Error) {
	var val []byte
	store := k.Store(ctx)
	iter := store.Iterator(nil, nil)
	if prefix != "" {
		iter = sdk.KVStorePrefixIterator(store, []byte(prefix))
	}

	for iter.Valid() {
		val = iter.Value()
		if len(val) > 1 {
			if !fn(val) {
				break
			}
		}
		iter.Next()
	}

	iter.Close()
	return
}

// Each calls `EachPrefix` with an empty prefix
func (k RecordKeeper) Each(ctx sdk.Context, fn func([]byte) bool) (err sdk.Error) {
	return k.EachPrefix(ctx, "", fn)
}

// IDKey returns the key for a given index
func (k RecordKeeper) IDKey(id uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", k.StorePrefix(), id))
}

// StorePrefix returns the root prefix of the key-value store
func (k RecordKeeper) StorePrefix() string {
	return fmt.Sprintf("%s:id:", k.StoreKey().Name())
}
