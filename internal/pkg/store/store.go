package store

import (
	"github.com/boltdb/bolt"
)

var defaultBucket = []byte("default")

type Store struct {
	db *bolt.DB
}

func NewStore(path string) (*Store, error) {

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}

	if err := store.CreateDefaultBucket(); err != nil {
		store.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) CreateDefaultBucket() error {
	return s.db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists(defaultBucket)
		return err
	})
}

func (s *Store) SetKey(key string, value []byte) error {
	return s.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket(defaultBucket)
		return b.Put([]byte(key), value)
	})
}

func (s *Store) GetKey(key string) ([]byte, error) {
	var result []byte

	err := s.db.View(func(t *bolt.Tx) error {
		b := t.Bucket(defaultBucket)
		result = b.Get([]byte(key))
		return nil
	})

	if err == nil {
		return result, nil
	}

	return nil, err

}

func (s *Store) Close() {
	s.db.Close()
}
