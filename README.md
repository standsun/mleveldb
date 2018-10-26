# mleveldb

mleveldb wrap the [syndtr/goleveldb](https://github.com/syndtr/goleveldb) which is a implementation of the [LevelDB key/value database](http:code.google.com/p/leveldb) to simple the leveldb operate.

# Install 

```
go get github.com/standsun/mleveldb
```

# Requirements

Need at least go1.5 or newer.

# Usage

```go

// init the leveldb
mleveldb.Init("./data")

db := mleveldb.New()

// set operate
db.Set("key_1", "value1")
db.Set("key_2", "value2")
...
db.Set("key_9", "value9")

// get operate
v, err := db.Get("key_1")
...


// find operate
r, err := db.Find("key_")
...

// range operate
r, err := db.Range("key_1", "key_5")
...

```
