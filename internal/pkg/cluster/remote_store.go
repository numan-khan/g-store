package cluster

import (
	"hash/fnv"
	"io/ioutil"
	"net/http"

	"github.com/numan-khan/g-store/internal/pkg/models"
)

type RemoteStore struct {
	config *models.Configuration
}

func NewRemoteStore(config *models.Configuration) (*RemoteStore, error) {
	return &RemoteStore{config: config}, nil
}

// func (s *RemoteStore) SetKey(key string, value []byte) error {
// 	destinationShard  := s.GetShard(key)
// 	rsp, err := http.Get("http://"+)
// }

func (s *RemoteStore) GetKey(key string) ([]byte, error) {
	destinationShard := s.GetShard(key)

	rsp, err := http.Get("http://" + destinationShard.Address + "get?key=" + key)
	if err != nil {
		return nil, err
	}

	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *RemoteStore) GetShard(key string) *models.Shard {
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

func (s *RemoteStore) Close() {
}
