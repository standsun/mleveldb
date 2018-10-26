package mleveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var (
	ldb *LevelDB
)

type Item struct {
	Key   string
	Value string
}

type LevelDB struct {
	db *leveldb.DB
}

func (l *LevelDB) Get(key string) (string, error) {
	value, err := l.db.Get([]byte(key), nil)
	return string(value), err
}

func (l *LevelDB) Set(key, value string) error {
	return l.db.Put([]byte(key), []byte(value), nil)
}

func (l *LevelDB) Delete(key string) error {
	return l.db.Delete([]byte(key), nil)
}

func (l *LevelDB) Find(prefix string) ([]*Item, error) {
	return l.each(l.db.NewIterator(util.BytesPrefix([]byte(prefix)), nil))
}

func (l *LevelDB) Range(start, end string) ([]*Item, error) {
	return l.each(l.db.NewIterator(&util.Range{
		Start: []byte(start),
		Limit: []byte(end),
	}, nil))
}

func (l *LevelDB) All() ([]*Item, error) {
	return l.each(l.db.NewIterator(nil, nil))
}

func (l *LevelDB) each(i iterator.Iterator) ([]*Item, error) {
	items := make([]*Item, 0, 10)

	for i.Next() {
		items = append(items, &Item{
			Key:   string(i.Key()),
			Value: string(i.Value()),
		})
	}

	i.Release()

	return items, i.Error()
}

func (l *LevelDB) Close() {
	l.db.Close()
}

func Init(path string) error {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}
	ldb = &LevelDB{
		db: db,
	}
	return nil
}

func New() *LevelDB {
	return ldb
}

func Close() {
	ldb.Close()
}
