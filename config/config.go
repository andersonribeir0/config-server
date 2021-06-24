package config

import (
	"encoding/json"
	"fmt"
	"github.com/andersonribeir0/config-server/logger"
	"os"
	"time"
)

type Settings struct {
	ConsulUrl           string
	Prefix              string
	AppName             string
	AutoRefresh         bool
	AutoRefreshSeconds  time.Duration
}

type Config struct {
	Keys   map[string]string
	Logger *logger.Log
}

var config *Config

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}

func getConfig(settings Settings) *Config {
	if config == nil {
		config = &Config{
			Keys:   nil,
			Logger: logger.NewLogger(settings.AppName),
		}
	}
	return config
}

func GetConfigKV() map[string]string{
	newKV := make(map[string]string)
	for k,v := range config.Keys {
		newKV[k] = v
	}
	return newKV
}

func Build(settings Settings, keys []string) *Config {
	var allKeys []string

	conf := getConfig(settings)
	log := conf.Logger
	conf.Keys = make(map[string]string)

	consulClient, err := GetConsulClient(settings)
	if err != nil {
		conf.Logger.Error("[CONFIG] It was not possible to communicate with consul cluster", err)
	}

	consulKeys, err := GetConsulKeys(consulClient, settings)
	if err != nil {
		conf.Logger.Error("[CONFIG] It was not possible to get consul keys", err)
	}

	allKeys = union(consulKeys, keys)

	//Must get local environment variable given a key if consul hit fails
	for _, key := range allKeys {
		pair, err := GetConsulKV(consulClient, settings, key)
		if err != nil {
			log.Error("[CONFIG] Impossible to get key value in consul", err)
			env := os.Getenv(key)
			conf.Keys[key] = env
			log.Info(fmt.Sprintf("[CONFIG] Getting key value %s from os env", env))
		} else if pair == nil {
			log.Info(fmt.Sprintf("[CONFIG] There is no such key named '%s' in consul", key))
		} else {
			result := conf.byteToString(pair.Value)
			conf.Keys[key] = result
		}
	}

	if consulNodes, err := GetConsulNodes(consulClient); err != nil {
		for k, _ := range consulNodes {
			conf.Keys[k] = consulNodes[k]
		}
	}

	conf.Dumps(settings)

	if settings.AutoRefresh {
		if settings.AutoRefreshSeconds*time.Second < time.Second {
			settings.AutoRefreshSeconds = 1
		}
		time.AfterFunc(settings.AutoRefreshSeconds*time.Second, func() {
			go Build(settings, keys)
		})
	}

	return conf
}

func (c Config) byteToString(bs []byte) string {
	ba := make([]byte, 0, len(bs))
	for _, b := range bs {
		ba = append(ba, b)
	}
	return string(ba)
}

func (c Config) Dumps(settings Settings) {
	data, err := json.Marshal(c.Keys)
	config := getConfig(settings)
	if err != nil {
		return
	}
	config.Logger.Info(fmt.Sprintf("[CONFIG] %s", string(data)))
}
