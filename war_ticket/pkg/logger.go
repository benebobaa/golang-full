package pkg

import (
	"github.com/rs/zerolog/log"
)

type LogFormat struct {
	HttpStatus  int
	ProcessTime uint
	Data        interface{}
	Error       string
	Type        string
	Message     string `json:"-"`
	IsSuccess   bool
}

func (l *LogFormat) Execute() {
	if l.IsSuccess {
		log.Info().Interface("log", l).Msg(l.Message)
	} else {
		log.Error().Interface("log", l).Msg(l.Message)
	}
}
