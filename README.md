![logo](./logo.jpg)

Iterable type keepers for [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) that implement [Active Record](https://en.wikipedia.org/wiki/Active_record_pattern).

There are currently two types of record keepers:
* RecordKeeper - An auto-incrementing `uint64` indexed keeper
- StringRecordKeeper - A `string` indexed keeper

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

### Deleting
```go
keeper.Delete(ctx, id)
```

### Updating
```go
updatedRecord := Record{}
keeper.Update(ctx, id, updatedRecord)
```

## String Example

```go
// setter
record := Record{}
keeper.Set(ctx, "key1", record)

// getter
keeper.Get(ctx, "key1", &record)
```
