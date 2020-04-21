package kits

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"os"
	"strings"
)

type logFormatter struct {
	log.TextFormatter
}

func (f *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	// This whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 31 // Gray
	case log.WarnLevel:
		levelColor = 33 // Yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // Red Color
	default:
		levelColor = 36 // Blue Color
	}
	return []byte(fmt.Sprintf("[%s] | \x1b[%dm%s\x1b[0m | %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

func InitLogger() {
	log.SetFormatter(&logFormatter{
		log.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
		},
	})

	log.SetOutput(os.Stderr)

	if cast.ToBool(os.Getenv("DEBUG")) {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}
