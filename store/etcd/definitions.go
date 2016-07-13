package etcdstore

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bobbytables/gangway/data"
	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
)

var (
	// GangwayKeyPrefix is the etcd key prefix
	GangwayKeyPrefix = "/gangway"

	// GangwayDefinitionsKey is the etcd key for all definitions
	GangwayDefinitionsKey = fmt.Sprintf("%s/%s", GangwayKeyPrefix, "definitions")
)

// RetrieveDefinitions implements store.Store
func (s *Store) RetrieveDefinitions() ([]data.Definition, error) {
	kapp := s.newKeysAPIFactory(s.etcdClient)
	resp, err := kapp.Get(context.TODO(), GangwayDefinitionsKey, &client.GetOptions{Recursive: true})
	if err != nil {
		return nil, err
	}

	var ds []data.Definition
	for _, node := range resp.Node.Nodes {
		label := strings.TrimPrefix(node.Key, GangwayDefinitionsKey+"/")
		jsonV := node.Value
		d := data.Definition{Label: label}

		if err := json.NewDecoder(strings.NewReader(jsonV)).Decode(&d); err != nil {
			return nil, err
		}

		ds = append(ds, d)
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
