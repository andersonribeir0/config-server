package config

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"strings"
)

func GetConsulClient(settings Settings) (*consulApi.Client, error) {
	consulConfig := consulApi.Config{
		Address: settings.ConsulUrl,
	}

	client, err := consulApi.NewClient(&consulConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetConsulKeys(client *consulApi.Client, settings Settings) ([]string, error) {
	var keys []string

	consulKeys, _, err := client.KV().Keys(
		fmt.Sprintf("%s/%s", settings.Prefix, settings.AppName),
		"",
		nil)
	if err != nil {
		return nil, err
	}

	for i := range consulKeys {
		if consulKeys[i] != "" {
			consulKeyChunks := strings.Split(consulKeys[i], "/")
			keys = append(keys, consulKeyChunks[len(consulKeyChunks)-1])
		}
	}
	return keys, nil
}

func GetConsulNodes(client *consulApi.Client) (map[string]string, error) {
	nodes, _, err := client.Catalog().Nodes(nil)
	consulNodes := make(map[string]string)
	if err != nil {
		return nil, err
	} else {
		for i := range nodes {
			node := nodes[i]
			if node != nil {
				consulNodes[nodes[i].Node] = nodes[i].Address
			}
		}
	}
	return consulNodes, nil
}

func GetConsulKV(client *consulApi.Client, settings Settings, key string) (*consulApi.KVPair, error) {
	keyWith := fmt.Sprintf("%s/%s/%s", settings.Prefix, settings.AppName, key)
	pair, _, err := client.KV().Get(keyWith, nil)
	return pair, err
}
