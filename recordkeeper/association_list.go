package recordkeeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AssociationList is a foreign key association between two stores
// associatedStoreKey:id:[associatedID]:storeKey:id:[ID]: -> [ID]
type AssociationList interface {
	Add(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64)
	Map(ctx sdk.Context, key sdk.StoreKey, id uint64, fn func(uint64))
	ReverseMap(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64))
}

// interface conformance check
var _ AssociationList = StringRecordKeeper{}

// Add adds a new association pair
func (k StringRecordKeeper) Add(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64) {
	association := fmt.Sprintf(
		"%s:id:%d:%s:id:%d",
		associatedKey.Name(), associatedID,
		key.Name(), id,
	)

	k.Set(ctx, association, id)
}

// Map iterates through associated ids and peforms function `fn`
func (k StringRecordKeeper) Map(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:id:%d:", associatedKey.Name(), associatedID)
	prefixBytes := []byte(prefix)

	iter := sdk.KVStorePrefixIterator(k.store(ctx), prefixBytes)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.codec.MustUnmarshalBinaryBare(iter.Value(), &id)
		fn(id)
	}
}

// ReverseMap reverse iterates through associated ids and peforms function `fn`
func (k StringRecordKeeper) ReverseMap(ctx sdk.Context, associatedKey sdk.StoreKey, associatedID uint64, fn func(uint64)) {
	prefix := fmt.Sprintf("%s:id:%d:", associatedKey.Name(), associatedID)
	prefixBytes := []byte(prefix)

	iter := sdk.KVStoreReversePrefixIterator(k.store(ctx), prefixBytes)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var id uint64
		k.codec.MustUnmarshalBinaryBare(iter.Value(), &id)
		fn(id)
	}
}
