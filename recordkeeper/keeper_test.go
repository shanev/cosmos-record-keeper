package recordkeeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/store"
	// "github.com/stretchr/testify/assert"
)

type Record struct {
	ID int64
}

func TestAddGetStory(t *testing.T) {
	// _, keeper := mockDB()

	// record := Record{
	// 	ID:         1,
	// }
}

func mockDB() (sdk.Context, RecordKeeper) {
	db := dbm.NewMemDB()

	storeKey := sdk.NewKVStoreKey("records")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()

	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())

	// codec := amino.NewCodec()
	// cryptoAmino.RegisterAmino(codec)
	// RegisterAmino(codec)

	keeper := NewRecordKeeper(storeKey)

	return ctx, keeper
}
