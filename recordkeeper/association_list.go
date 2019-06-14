package recordkeeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// interface conformance check
var _ Uint64AssociationKeeper = RecordKeeper{}
var _ AddressAssociationKeeper = RecordKeeper{}

// Uint64AssociationKeeper is a foreign key association between two stores
// associatedStoreKey:id:[associatedID]:storeKey:id:[ID]: -> [ID]
type Uint64AssociationKeeper interface {
	Push(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64)
	Map(ctx sdk.Context, associatedKey sdk.StoreKey, id uint64, fn func(uint64))
	ReverseMap(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64))
}

type AddressAssociationKeeper interface {
	PushWithAddress(ctx sdk.Context, key, associatedKey sdk.StoreKey, id uint64, address sdk.AccAddress)
	MapByAddress(ctx sdk.Context, associatedKey sdk.StoreKey, address sdk.AccAddress, fn func(uint64))
	ReverseMapByAddress(ctx sdk.Context, associatedKey sdk.StoreKey, address sdk.AccAddress, fn func(uint64))
}

// Push adds a new association pair
func (k RecordKeeper) Push(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64) {
	association := fmt.Sprintf(
		"%s:id:%d:%s:id:%d",
		associatedKey.Name(), associatedID,
		key.Name(), id,
	)

	k.stringSetBare(ctx, association, id)
}

// Map iterates through associated ids and performs function `fn`
func (k RecordKeeper) Map(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:id:%d:", associatedKey.Name(), associatedID)
	prefixBytes := []byte(prefix)
	k.iterate(ctx, prefixBytes, fn, false)
}

// ReverseMap reverse iterates through associated ids and performs function `fn`
func (k RecordKeeper) ReverseMap(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:id:%d:", associatedKey.Name(), associatedID)
	prefixBytes := []byte(prefix)
	k.iterate(ctx, prefixBytes, fn, true)
}

// PushWithAddress adds a new association pair
func (k RecordKeeper) PushWithAddress(ctx sdk.Context, key, associatedKey sdk.StoreKey, id uint64, address sdk.AccAddress) {
	association := fmt.Sprintf(
		"%s:address:%s:%s:id:%d",
		associatedKey.Name(), address.String(),
		key.Name(), id,
	)
	k.stringSetBare(ctx, association, id)
}

// MapByAddress iterates through associated ids and performs function `fn`
func (k RecordKeeper) MapByAddress(ctx sdk.Context, associatedKey sdk.StoreKey, address sdk.AccAddress, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:address:%s:", associatedKey.Name(), address.String())
	prefixBytes := []byte(prefix)
	k.iterate(ctx, prefixBytes, fn, false)
}

// ReverseMapByAddress reverse iterates through associated ids and performs function `fn`
func (k RecordKeeper) ReverseMapByAddress(ctx sdk.Context, associatedKey sdk.StoreKey, address sdk.AccAddress, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:address:%s:", associatedKey.Name(), address.String())
	prefixBytes := []byte(prefix)
	k.iterate(ctx, prefixBytes, fn, true)
}

func (k RecordKeeper) stringSetBare(ctx sdk.Context, key string, value interface{}) {
	valueBytes := k.codec.MustMarshalBinaryBare(value)
	k.stringSetBytes(ctx, key, valueBytes)
}


func (k RecordKeeper) iterate(ctx sdk.Context, prefix []byte, fn func(uint64), reverse bool){
	iter := sdk.KVStorePrefixIterator(k.store(ctx), prefix)
	if reverse {
		iter = sdk.KVStoreReversePrefixIterator(k.store(ctx), prefix)
	}
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.codec.MustUnmarshalBinaryBare(iter.Value(), &id)
		fn(id)
	}
}
