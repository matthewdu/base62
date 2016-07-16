# base62

[![CircleCI](https://circleci.com/gh/matthewdu/base62.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/matthewdu/base62)

Convert to and from `uint64` into base62 `string`. Useful for shortening long integer IDs.
I personally use for Google Cloud Datastore IDs.

### Install
```bash
go get -u github.com/matthewdu/base62
```

### Test (with benchmark)
```bash
go test -benchmem -bench .
```

### Use
```go
var n uint64 = 3781504209452600
s := base62.Encode(n) // hjNv8tS3K 
m, _ := base62.Decode(s) // 3781504209452600
```
