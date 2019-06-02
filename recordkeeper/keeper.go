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

// Set sets a key, value pair in the store
func (k RecordKeeper) Set(ctx sdk.Context, key uint64, value []byte) {
	idBytes := k.idKey(key)
	k.store(ctx).Set(idBytes, value)
}

// Get gets a value given a key
func (k RecordKeeper) Get(ctx sdk.Context, key uint64) []byte {
	idBytes := k.idKey(key)
	return k.store(ctx).Get(idBytes)
}

// NextID increments and returns the next available id by 1
func (k RecordKeeper) NextID(ctx sdk.Context) (id uint64) {
	idBytes := k.store(ctx).Get(k.lenKey())
	if idBytes == nil {
		initialIndex := uint64(1)
		k.setLen(ctx, initialIndex)

		return initialIndex
	}

	err := json.Unmarshal(idBytes, &id)
	if err != nil {
		panic(err)
	}

	nextID := id + 1
	k.setLen(ctx, nextID)

	return nextID
}

// EachPrefix calls `fn` for each record in a store with a given prefix. Iteration will stop if `fn` returns false
func (k RecordKeeper) EachPrefix(ctx sdk.Context, prefix string, fn func([]byte) bool) (err sdk.Error) {
	var val []byte
	store := k.store(ctx)
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

func (k RecordKeeper) idKey(id uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", k.storePrefix(), id))
}

func (k RecordKeeper) lenKey() []byte {
	return []byte(k.storeKey.Name() + ":len")
}

func (k RecordKeeper) setLen(ctx sdk.Context, len uint64) {
	idBytes, err := json.Marshal(len)
	if err != nil {
		panic(err)
	}
	k.store(ctx).Set(k.lenKey(), idBytes)
}

func (k RecordKeeper) store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

func (k RecordKeeper) storePrefix() string {
	return fmt.Sprintf("%s:id:", k.storeKey.Name())
}
