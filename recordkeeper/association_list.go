package recordkeeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AssociationList is a foreign key association between two stores
// assoc:storeKey:id:ID:associatedStoreKey:id:associatedID -> associatedID
type AssociationList interface {
	Add(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64)
	Map(ctx sdk.Context, key sdk.StoreKey, id uint64, fn func(uint64))
}

// interface conformance check
var _ AssociationList = StringRecordKeeper{}

// Add adds a new association pair
func (k StringRecordKeeper) Add(ctx sdk.Context, key, associatedKey sdk.StoreKey, id, associatedID uint64) {
	tuple := fmt.Sprintf(
		"assoc:%s:id:%d:%s:id:%d",
		key.Name(), id,
		associatedKey.Name(), associatedID)

	k.Set(ctx, tuple, associatedID)
}

// Map iterates through associated ids and peforms function `fn`
func (k StringRecordKeeper) Map(ctx sdk.Context, key sdk.StoreKey, id uint64, fn func(uint64)) {
	searchKey := fmt.Sprintf(
		"assoc:%s:id:%d:",
		key.Name(), id)

	prefix := []byte(searchKey)

	// iterates through keyspace to find all value ids
	iter := sdk.KVStorePrefixIterator(k.store(ctx), prefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var associatedID uint64
		k.codec.MustUnmarshalBinaryBare(iter.Value(), &associatedID)
		fn(associatedID)
	}
}
