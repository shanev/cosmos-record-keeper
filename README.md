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

### Initialization

```go
    storeKey := sdk.NewKVStoreKey("record")
    keeper := Keeper{
        &NewRecordKeeper(storeKey),
    }
```

### Adding

```go
    record := Record{}
    recordBytes, _ := json.Marshal(record)
    id := keeper.Add(ctx, recordBytes)
```

### Getting

```go
    recordBytes := keeper.Get(ctx, id)
    var r Record
    json.Unmarshal(recordBytes, &r)
```
