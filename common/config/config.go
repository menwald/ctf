package config

import (
	"flag"
	"fmt"

	"github.com/a8m/envsubst"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func LoadConfig() {
	configFile := flag.String("config-file", "./configs/config.yaml", "Config file location.")
	flag.Parse()
	v := viper.New()
	b, err := envsubst.ReadFileRestricted(configFile, true, false)
	if err != nil {

	}

	configVals, err := config.NewConfig(configFile)
	if err != nil {
		panic(fmt.Sprintf("failed to load config.  Error: %s", err.Error()))
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info().Msg("no configuration file found")
		} else {
			log.Err(err).Msg("unable to load configuration file")
		}
	}
	log.Debug().Str("config-file", viper.ConfigFileUsed()).Msg("loaded configuration file")
}
