package output

import (
	"encoding/json"
	"fmt"
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
	serialized := make([]byte, 0, 300)
	serialized = append(serialized, []byte(entry.Time.Format("2006-01-02 15:04:05.000 "))...)
	level := fmt.Sprintf("%-19s",colorizeLevel(entry.Level))
	serialized = append(serialized, []byte(level)...)
	serialized = append(serialized, []byte(entry.Message)...)
	if data, err := json.Marshal(entry.Data); err == nil && len(data) > 2{
		serialized = append(serialized, ' ')
		serialized = append(serialized, data...)
	}
	return append(serialized, '\n'), nil
}

// colorizeLevels applies a color based to the log.Level
func colorizeLevel(l log.Level) string{
	switch l {
	case log.ErrorLevel:
		return fmt.Sprintf("%-10s","[" + Red + strings.ToUpper(l.String()) + Reset + "] ")
	case log.WarnLevel:
		return fmt.Sprintf("%-10s","[" + Yellow + strings.ToUpper(l.String()) + Reset + "] ")
	case log.InfoLevel:
		return fmt.Sprintf("%-10s","[" + Green + strings.ToUpper(l.String()) + Reset + "] ")
	case log.DebugLevel:
		return fmt.Sprintf("%-10s","[" +  Gray + strings.ToUpper(l.String()) + Reset + "] ")
	case log.TraceLevel:
		return fmt.Sprintf("%-10s","[" + DarkGray + strings.ToUpper(l.String()) + Reset + "] ")
	default:
		return fmt.Sprintf("%-10s","[" + strings.ToUpper(l.String()) + "] ")
	}
}