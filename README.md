![logo](./logo.jpg)

A `uint64` indexed, iterable type keeper for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) that implements the [Active Record pattern](https://en.wikipedia.org/wiki/Active_record_pattern).

[![CircleCI](https://circleci.com/gh/shanev/cosmos-record-keeper.svg?style=svg)](https://circleci.com/gh/shanev/cosmos-record-keeper)

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
var record Record
keeper.Get(ctx, id, &record)
```

### Iterating

```go
keeper.Each(ctx, func(recordBytes []byte) bool {
    var r Record
    keeper.codec.MustUnmarshalBinaryLengthPrefixed(recordBytes, &r)
    // do something with `Record` r
    return true
})
```

## API

```go
Add(ctx sdk.Context, value interface{}) uint64
Delete(ctx sdk.Context, id uint64, t interface{})
EachPrefix(ctx sdk.Context, prefix string, fn func([]byte) bool) (err sdk.Error)
Each(ctx sdk.Context, fn func([]byte) bool) (err sdk.Error)
Get(ctx sdk.Context, key uint64, value interface{}) sdk.Error
IncrementID(ctx sdk.Context) (id uint64)
Set(ctx sdk.Context, key uint64, value []byte)
Update(ctx sdk.Context, key uint64, value interface{}) uint64
```
