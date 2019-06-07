![logo](./logo.jpg)

Keeper utilities for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) that attempt to implement [Active Record](https://en.wikipedia.org/wiki/Active_record_pattern).

Currently, three interfaces are supported:
- `Uint64KeyedIterableKeeper` - An auto-incrementing `uint64` indexed keeper
- `StringKeyedKeeper` - A `string` indexed keeper
- `Uint64AssociationKeeper` - One-to-many associations with another store

[![CircleCI](https://circleci.com/gh/shanev/cosmos-record-keeper.svg?style=svg)](https://circleci.com/gh/shanev/cosmos-record-keeper)
[![Go Report Card](https://goreportcard.com/badge/github.com/shanev/cosmos-record-keeper)](https://goreportcard.com/report/github.com/shanev/cosmos-record-keeper)

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
    RecordKeeper(storeKey, codec),
}
```

### Adding

```go
record := Record{}
id := k.Add(ctx, record)
```

### Getting

```go
var record Record
k.Get(ctx, id, &record)
```

### Iterating

```go
k.Each(ctx, func(recordBytes []byte) bool {
    var r Record
    k.codec.MustUnmarshalBinaryLengthPrefixed(recordBytes, &r)
    // do something with `Record` r
    return true
})
```

### Deleting
```go
k.Delete(ctx, id)
```

### Updating
```go
updatedRecord := Record{}
k.Update(ctx, id, updatedRecord)
```

## String Example

```go
// setter
record := Record{}
k.StringSet(ctx, "key1", record)

// getter
k.StringGet(ctx, "key1", &record)
```

## One-to-many Association Example

```go
k.Push(ctx, k.StoreKey, k2.StoreKey, uint64(1), uint64(2))
k.Map(ctx, k2.StoreKey, uint64(2), func(id uint64) {
    // id == 1
})
```
