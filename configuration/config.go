package config

import (
	file_manager "github.com/udayangaac/start-and-go/lib/file-manager"
)

type Configuration interface {
	Read(fm file_manager.FileManager)
}
type Configurations []Configuration

func (configs Configurations) Init(fm file_manager.FileManager) {
	for _, c := range configs {
		c.Read(fm)
	}
}
