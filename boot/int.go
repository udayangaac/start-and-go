package boot

import (
	log "github.com/sirupsen/logrus"
	config "github.com/udayangaac/start-and-go/configuration"
	file_manager "github.com/udayangaac/start-and-go/lib/file-manager"
	"os"
)

func Init() {
	// initialize the logger
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	// initialize the yaml file reader
	yamlReader := file_manager.NewYamlManager()

	// get env related config
	config.Configurations{
		new(config.ServerConfig),
	}.Init(yamlReader)

	config.Configurations{
		new(config.ServerConfig),
		new(config.DatabaseConfig),
	}.Init(yamlReader)
}
