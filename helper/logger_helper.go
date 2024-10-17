package helper

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Logger *log.Logger

func LoggerHelper() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
}
