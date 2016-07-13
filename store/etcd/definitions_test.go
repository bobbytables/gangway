package etcdstore

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/bobbytables/gangway/data"
	"github.com/coreos/etcd/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type FakeKeysAPI struct {
	mock.Mock
}

func (f *FakeKeysAPI) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	o := f.Called(ctx, key, opts)
	return o.Get(0).(*client.Response), o.Error(1)
}

func (f *FakeKeysAPI) Set(ctx context.Context, key, value string, opts *client.SetOptions) (*client.Response, error) {
	return nil, nil
}
func (f *FakeKeysAPI) Delete(ctx context.Context, key string, opts *client.DeleteOptions) (*client.Response, error) {
	return nil, nil
}
func (f *FakeKeysAPI) Create(ctx context.Context, key, value string) (*client.Response, error) {
	return nil, nil
}
func (f *FakeKeysAPI) CreateInOrder(ctx context.Context, dir, value string, opts *client.CreateInOrderOptions) (*client.Response, error) {
	return nil, nil
}
func (f *FakeKeysAPI) Update(ctx context.Context, key, value string) (*client.Response, error) {
	return nil, nil
}

func (f *FakeKeysAPI) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	return nil
}

func makeStore() (*Store, *FakeKeysAPI) {
	keysAPI := &FakeKeysAPI{}
	keysFactory := func(c client.Client) client.KeysAPI {
		return keysAPI
	}

	s := &Store{
		newKeysAPIFactory: keysFactory,
	}

	return s, keysAPI
}

func TestRetrieveDefinitions(t *testing.T) {
	s, keysAPI := makeStore()

	resp := &client.Response{
		Node: &client.Node{
			Nodes: []*client.Node{{Key: "chicken-cat", Value: `{"source": "hello-world", "dockerfile": "mydockerfile", "tag":"sometag"}`}},
		},
	}

	keysAPI.On("Get",
		context.TODO(),
		GangwayDefinitionsKey,
		&client.GetOptions{Recursive: true},
	).Return(resp, nil)

	ds, err := s.RetrieveDefinitions()

	require.Nil(t, err)
	assert.Len(t, ds, 1)
	d := data.Definition{Label: "chicken-cat", Source: "hello-world", Dockerfile: "mydockerfile", Environment: nil, Tag: "sometag"}
	assert.Equal(t, d, ds[0], "definition matches etcd return value")
}
