package pkg

import (
	"github.com/rs/zerolog/log"
)

type LogFormat struct {
	HttpStatus  int
	ProcessTime uint
	Data        interface{}
	Error       string
	Status      string
	Message     string
	IsSuccess   bool
}

func (l *LogFormat) Execute() {
	if l.IsSuccess {
		log.Info().Interface("log", l).Msg(l.Message)
	} else {
		log.Error().Interface("log", l).Msg(l.Message)
	}
}
