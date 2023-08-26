package main

import (
	"github.com/menwald/ctf/common"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	common.LoadConfig()
}
