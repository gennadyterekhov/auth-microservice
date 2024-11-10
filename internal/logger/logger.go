package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Debugln(msg ...interface{}) {
	log.Debug().Msg(makeMessage(msg))
}

func Errorln(msg ...string) {
	log.Error().Msg(makeMessage(msg))
}

func makeMessage(msg ...interface{}) string {
	fullMessage := ""
	for _, m := range msg {
		fullMessage += fmt.Sprintf("%+v ", m)
	}
	return fullMessage
}
