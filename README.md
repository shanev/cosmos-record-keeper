# Cosmos Record Keeper

A `uint64` indexed iterable type keeper for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

## Example

Embed a `RecordKeeper` struct inside a `Keeper`.

i.e:
```go
type Keeper struct {
    RecordKeeper
}
```

Define an index field for a type:

```go
type Record struct {
    ID uint64
}
```

### Initialization

```go
    storeKey := sdk.NewKVStoreKey("records")
    keeper := Keeper{&NewRecordKeeper(storeKey)}
```

### Setting

```go
    record := Record{
        ID: keeper.NextID(ctx),
    }
    recordBytes, _ := json.Marshal(record)
    keeper.Set(ctx, record.ID, recordBytes)
```

### Getting

```go
    recordBytes := keeper.Get(ctx, record.ID)
    var r Record
    json.Unmarshal(recordBytes, &r)
```
