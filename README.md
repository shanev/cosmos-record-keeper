![logo](./logo.jpg)

A `uint64` indexed, iterable type keeper for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) that implements the [Active Record pattern](https://en.wikipedia.org/wiki/Active_record_pattern).

## Getting Started

### Install library

```
go get github.com/shanev/cosmos-record-keeper
```

## Example

Embed a `RecordKeeper` struct inside a `Keeper`.

```go
type Keeper struct {
    RecordKeeper
}
```

### Initialization

```go
keeper := Keeper{
    &NewRecordKeeper(storeKey, codec),
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