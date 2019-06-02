package recordkeeper

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

type Record struct{}

func TestRecord(t *testing.T) {
	ctx, keeper := mockDB()
	assert.NotNil(t, ctx)
	assert.NotNil(t, keeper)

	record := Record{}

	// marshal record
	recordBytes, err := json.Marshal(record)
	assert.NoError(t, err)

	// add to the store
	id := keeper.Add(ctx, recordBytes)

	// getting
	expectedRecordBytes := keeper.Get(ctx, id)
	assert.Equal(t, expectedRecordBytes, recordBytes)

	// test iteration
	err = keeper.Each(ctx, func(recordBytes []byte) bool {
		var r Record
		err := json.Unmarshal(recordBytes, &r)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), id)
		return true
	})
}

func mockDB() (sdk.Context, RecordKeeper) {
	db := dbm.NewMemDB()

	storeKey := sdk.NewKVStoreKey("records")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())

	keeper := NewRecordKeeper(storeKey)

	return ctx, keeper
}
