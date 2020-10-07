package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	file_manager "github.com/udayangaac/start-and-go/lib/file-manager"
)

var ServerConf ServerConfig

type ServerConfig struct {
	Port           int    `yaml:"port"`
	Jwt            JWT    `yaml:"jwt"`
	MetricsPort    int    `yaml:"metrics_port"`
	ResourcePath   string `yaml:"resource_path"`
	FileServerPort int    `yaml:"file_server_port"`
}

type JWT struct {
	Key      string `yaml:"key"`
	Duration int    `yaml:"duration"`
}

func (sc *ServerConfig) Read(fm file_manager.FileManager) {
	path := fmt.Sprintf(`config/server.yaml`)
	err := fm.Read(path, &ServerConf)
	if err != nil {
		log.Fatal("Unable to read the server.yaml file.")
	}
}
