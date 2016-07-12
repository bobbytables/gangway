package etcdstore

import (
	"time"

	"github.com/bobbytables/gangway/store"

	"github.com/coreos/etcd/client"
)

var _ client.Client

// Store implements store.Store for interacting with etcd
type Store struct {
	etcdClient client.Client
}

var _ store.Store = &Store{}

// NewStore starts a storage interface against etcd endpoints
func NewStore(endpoints []string) (*Store, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Store{
		etcdClient: c,
	}, nil
}
