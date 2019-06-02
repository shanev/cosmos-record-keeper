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

	// add to the store
	record := Record{}
	id := keeper.Add(ctx, record)

	// getting
	expectedRecord := keeper.Get(ctx, id)
	assert.Equal(t, expectedRecord, record)

	// test iteration
	keeper.Each(ctx, func(recordBytes []byte) bool {
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
