package cluster

import (
	"hash/fnv"

	"github.com/numan-khan/g-store/internal/pkg/models"
	"github.com/numan-khan/g-store/internal/pkg/store"
)

type DistributedStore struct {
	localStore  *store.Store
	remoteStore *RemoteStore
	config      *models.Configuration
	shard       string
}

func NewDistributedStore(path string, shard string, config *models.Configuration) (*DistributedStore, error) {
	store, err := store.NewStore(path)
	if err != nil {
		return nil, err
	}

	rStore, err := NewRemoteStore(config)

	return &DistributedStore{localStore: store, remoteStore: rStore, config: config, shard: shard}, nil
}

func (s *DistributedStore) SetKey(key string, value []byte) error {

	if sh := s.GetShard(key); sh.Name == s.shard {
		return s.localStore.SetKey(key, value)
	} else {
		return s.remoteStore.SetKey(key, value)
	}
}

func (s *DistributedStore) GetKey(key string) ([]byte, error) {

	if sh := s.GetShard(key); sh.Name == s.shard {
		return s.localStore.GetKey(key)
	} else {
		return s.remoteStore.GetKey(key)
	}
}

func (s *DistributedStore) GetShard(key string) *models.Shard {
	fn := fnv.New64()
	fn.Write([]byte(key))
	shardIdx := int(fn.Sum64() % uint64(len(s.config.Shards)))
	var shard *models.Shard
	for _, s := range s.config.Shards {
		if s.Idx == shardIdx {
			shard = &s
		}
	}

	return shard
}

func (s *DistributedStore) Close() {
	s.localStore.Close()
}
