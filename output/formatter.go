package output

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"
)

type TextFormatter struct {
}

func (f *TextFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Note this doesn't include Time, Level and Message which are available on
	// the Entry. Consult `godoc` on information about those fields or read the
	// source of the official loggers.
	//var err error
	serialized := make([]byte, 1, 1000)
	serialized[0] = byte('[')
	l := []byte(strings.ToUpper(entry.Level.String()))
	serialized = append(serialized, l...)
	serialized = append(serialized, ']', ' ')
	serialized = append(serialized, []byte(entry.Time.Format("2006-01-02T15:04:05.000 "))...)
	serialized = append(serialized, '-', ' ')
	serialized = append(serialized, []byte(entry.Message)...)
	if data, err := json.Marshal(entry.Data); err == nil && len(data) > 2{
		serialized = append(serialized, ' ')
		serialized = append(serialized, data...)
	}
	return append(serialized, '\n'), nil
}