package etcdstore

import (
	"github.com/bobbytables/gangway/data"
	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
)

// RetrieveDefinitions implements store.Store
func (s *Store) RetrieveDefinitions() ([]data.Definition, error) {
	kapp := client.NewKeysAPI(s.etcdClient)
	resp, err := kapp.Get(context.TODO(), "gangway/definitions", &client.GetOptions{Recursive: true})
	if err != nil {
		return nil, err
	}

	var ds []data.Definition
	for _, node := range resp.Node.Nodes {
		m := getMapFromNode(node)
		ds = append(ds, data.Definition{Label: node.Key, Source: m["/gangway/definitions/default/source"]})
	}

	return ds, nil
}

// AddDefinition implements store.Store
func (s *Store) AddDefinition(d data.Definition) error {
	return nil
}

func getMapFromNode(n *client.Node) map[string]string {
	if !n.Dir {
		return nil
	}

	m := make(map[string]string)
	for _, n := range n.Nodes {
		if n.Dir {
			continue
		}

		m[n.Key] = n.Value
	}

	return m
}
