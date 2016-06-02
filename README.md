# base62

Convert to and from int64 into base-62 string. Useful for shortening long integer IDs.
I personally use for Google Cloud Datastore IDs.

```golang
var n int64 = 3781504209452600
s := base62.Encode(n) // hjNv8tS3K 
m, _ := base62.Decode(s) // 3781504209452600
```
