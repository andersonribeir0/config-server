package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/andersonribeir0/config-server/config"
	"github.com/andersonribeir0/config-server/logger"
	"net/http"
	"time"
)

type Server struct {
	HttpPort         		 string
	ConsulURL                string
	ConsulPort   		     string
	ConsulPrefix 		     string
	ConsulAutoRefresh        bool
	ConsulAutoRefreshSeconds int64
	AppName      		     string
}

var serverConfig *config.Config
var log *logger.Log

func (s Server) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(config.GetConfigKV())
	w.Header().Add("Content-Type", "application/json")
	if err == nil {
		_, err = w.Write(b)
	} else {
		log.Error(fmt.Sprintf("An error occurred when getting config KV: %s", err.Error()), err)
		return
	}
}

func (s Server) Serve() {
	log = logger.NewLogger(s.AppName)
	serverConfig = config.Start(config.Settings{
		ConsulUrl:   s.ConsulURL + s.ConsulPort,
		Prefix:      s.ConsulPrefix,
		AppName:     s.AppName,
		AutoRefresh: s.ConsulAutoRefresh,
		AutoRefreshSeconds: time.Duration(s.ConsulAutoRefreshSeconds),
	}, nil)

	http.HandleFunc("/", s.ConfigHandler)
	if err := http.ListenAndServe(s.HttpPort, nil); err != nil {
		log.Error(fmt.Sprintf("An error occurred when trying to run ConfigHandler: %s", err.Error()), err)
	}
}
