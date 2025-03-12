package utils

import (
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/config/file"
	"github.com/dobyte/due/v2/etc"
)

func SetEtcConfigurator(path string) {
	etc.SetConfigurator(
		config.NewConfigurator(
			config.WithSources(
				file.NewSource(file.WithPath(path)),
			),
		),
	)
}
