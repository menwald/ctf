package gamemap

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/yanndr/tmx"
)

func JustOpenIt() error {
	cwdp, err := os.Getwd()
	if err != nil {
		log.Err(err).Send()
	}
	fmt.Println(cwdp)
	path := "common/resources/"
	f, err := os.Open(path + "testinggrassfield.tmx")
	if err != nil {
		log.Err(err).Send()
		return err
	}
	defer f.Close()
	m, err := tmx.Parse(f)
	if err != nil {
		log.Err(err).Send()
		return err
	}
	log.Info().Int("width", m.Width).Msg("loaded tmx")
	return nil
}
