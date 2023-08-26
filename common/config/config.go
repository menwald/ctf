package config

import (
	"bytes"
	"flag"
	"path/filepath"
	"strings"

	"github.com/a8m/envsubst"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	filename string
	viper    *viper.Viper
}

func (c Config) Filename() string {
	return c.filename
}

func (c Config) GetString(key string) string {
	return c.viper.GetString(key)
}

func (c Config) GetInt(key string) int {
	return c.viper.GetInt(key)
}

func Load() (Config, error) {
	configFile := flag.String("config-file", "./configs/config.yaml", "Config file location.")
	flag.Parse()
	ret := Config{filename: *configFile, viper: viper.New()}
	ret.viper.SetConfigType("yaml")
	b, err := envsubst.ReadFileRestricted(*configFile, true, false)
	if err != nil {
		log.Err(err).Msg("unable to read config file")
		return ret, err
	}

	ret.viper.SetConfigType(strings.TrimPrefix(filepath.Ext(*configFile), "."))
	if err := ret.viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Err(err).Msg(("unable to load config file"))
			return ret, err
		}
	}

	return ret, nil
}
