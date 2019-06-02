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

type Record struct {
	ID uint64
}

func TestRecord(t *testing.T) {
	ctx, keeper := mockDB()
	assert.NotNil(t, ctx)
	assert.NotNil(t, keeper)

	record := Record{
		ID: keeper.NextID(ctx),
	}
	assert.Equal(t, uint64(1), record.ID)

	// marshal record
	recordBytes, err := json.Marshal(record)
	assert.NoError(t, err)

	// set in kvstore
	keeper.Set(ctx, record.ID, recordBytes)

	// getting
	expectedRecordBytes := keeper.Get(ctx, record.ID)
	assert.Equal(t, expectedRecordBytes, recordBytes)

	// test iteration
	err = keeper.Each(ctx, func(recordBytes []byte) bool {
		var r Record
		err := json.Unmarshal(recordBytes, &r)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), r.ID)
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
