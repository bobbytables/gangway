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
		d, err := definitionFromNode(node)
		if err != nil {
			return nil, err
		}

		ds = append(ds, d)
	}

	return ds, nil
}

// AddDefinition implements store.Store
func (s *Store) AddDefinition(d data.Definition) error {
	kapp := s.newKeysAPIFactory(s.etcdClient)

	js, err := json.Marshal(d)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s/%s", GangwayDefinitionsKey, d.Label)
	_, err = kapp.Set(context.TODO(), key, string(js), nil)

	return err
}

func definitionFromNode(n *client.Node) (data.Definition, error) {
	label := strings.TrimPrefix(n.Key, GangwayDefinitionsKey+"/")
	jsonV := n.Value
	d := data.Definition{Label: label}

	if err := json.NewDecoder(strings.NewReader(jsonV)).Decode(&d); err != nil {
		return data.Definition{}, err
	}

	return d, nil
}
