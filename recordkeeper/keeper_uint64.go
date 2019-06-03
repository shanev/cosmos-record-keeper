package recordkeeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Uint64IterableKeeper defines methods for the active record pattern
type Uint64IterableKeeper interface {
	Add(ctx sdk.Context, value interface{}) uint64
	Delete(ctx sdk.Context, id uint64) uint64
	EachPrefix(ctx sdk.Context, prefix string, fn func([]byte) bool) (err sdk.Error)
	Each(ctx sdk.Context, fn func([]byte) bool) (err sdk.Error)
	Get(ctx sdk.Context, key uint64, value interface{}) sdk.Error
	IncrementID(ctx sdk.Context) uint64
	Set(ctx sdk.Context, key uint64, value interface{})
	Update(ctx sdk.Context, key uint64, value interface{}) uint64
}

// interface conformance check
var _ Uint64IterableKeeper = RecordKeeper{}

// RecordKeeper data type with a default codec
type RecordKeeper struct {
	StoreKey sdk.StoreKey
	Codec    *codec.Codec
}

// NewRecordKeeper creates a new record keeper for module keepers to embed
func NewRecordKeeper(storeKey sdk.StoreKey, codec *codec.Codec) RecordKeeper {
	return RecordKeeper{storeKey, codec}
}

// Add adds a value to the store
func (k RecordKeeper) Add(ctx sdk.Context, value interface{}) uint64 {
	id := k.IncrementID(ctx)
	k.Set(ctx, id, value)

	return id
}

// Delete deletes a value from the store
// NOTE: This retains the key, and sets the value to nil.
// Make sure to check for nil when getting values from the store.
func (k RecordKeeper) Delete(ctx sdk.Context, id uint64) uint64 {
	k.Set(ctx, id, nil)

	return id
}

// EachPrefix calls `fn` for each record in a store with a given prefix.
// Iteration will stop if `fn` returns false.
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

// Get gets a value given a key
func (k RecordKeeper) Get(ctx sdk.Context, key uint64, value interface{}) sdk.Error {
	idBytes := k.idKey(key)
	recordBytes := k.store(ctx).Get(idBytes)
	if recordBytes == nil {
		return sdk.ErrInternal("Value not found at index " + strconv.FormatUint(key, 10))
	}
	k.Codec.MustUnmarshalBinaryLengthPrefixed(recordBytes, value)

	return nil
}

// IncrementID increments the index and stores the new value
func (k RecordKeeper) IncrementID(ctx sdk.Context) (id uint64) {
	idBytes := k.store(ctx).Get(k.lenKey())
	if idBytes == nil {
		initialIndex := uint64(1)
		k.SetLen(ctx, initialIndex)
		return initialIndex
	}

	k.Codec.MustUnmarshalBinaryBare(idBytes, &id)
	nextID := id + 1
	k.SetLen(ctx, nextID)

	return nextID
}

// Set sets a key, value pair in the store
func (k RecordKeeper) Set(ctx sdk.Context, key uint64, value interface{}) {
	valueBytes := k.Codec.MustMarshalBinaryLengthPrefixed(value)
	k.setBytes(ctx, key, valueBytes)
}

// SetLen sets the value of len for the store
func (k RecordKeeper) SetLen(ctx sdk.Context, len uint64) {
	idBytes := k.Codec.MustMarshalBinaryBare(len)
	k.store(ctx).Set(k.lenKey(), idBytes)
}

// Update updates a value to the store
func (k RecordKeeper) Update(ctx sdk.Context, key uint64, value interface{}) uint64 {
	k.Set(ctx, key, value)

	return key
}

// Internal

func (k RecordKeeper) idKey(id uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", k.storePrefix(), id))
}

func (k RecordKeeper) lenKey() []byte {
	return []byte(k.StoreKey.Name() + ":len")
}

func (k RecordKeeper) setBytes(ctx sdk.Context, key uint64, value []byte) {
	idBytes := k.idKey(key)
	k.store(ctx).Set(idBytes, value)
}

func (k RecordKeeper) store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.StoreKey)
}

func (k RecordKeeper) storePrefix() string {
	return fmt.Sprintf("%s:id:", k.StoreKey.Name())
}
