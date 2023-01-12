package main

import (
	"WEBAPP/apiserver"
	"flag"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configurator-path", "apiserver/config/config.yaml", "path to configurator file")
}

func main() {
	flag.Parse()
	logrus.Trace("config path: ", configPath)
	server, err := apiserver.New().ConfigureServer(configPath)
	if err != nil {
		logrus.Error(err)
	}
	server.Start()
}
