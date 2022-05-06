package models

import "errors"

type Configuration struct {
	Shards []Shard
}

func (c *Configuration) GetShardByName(name string) (*Shard, error) {

	for _, s := range c.Shards {
		if s.Name == name {
			return &s, nil
		}
	}

	return nil, errors.New("can't found shard with provided name")
}
