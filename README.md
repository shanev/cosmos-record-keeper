![logo](./logo.jpg)

A `uint64` indexed, iterable type keeper for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

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
    id := keeper.Add(ctx, record)
```

### Getting

```go
    record := keeper.Get(ctx, id)
```

### Iterating

```go
    keeper.Each(ctx, func(recordBytes []byte) bool {
        var r Record
        json.Unmarshal(recordBytes, &r)
        // do something with `Record` r
        return true
    })
```