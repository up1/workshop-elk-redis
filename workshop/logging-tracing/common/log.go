package common

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
  id      int
  message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
  *logrus.Logger
  name string
}

// NewLogger initializes the standard logger
func NewLogger(name string) *StandardLogger {
	var baseLogger = logrus.New()
	var standardLogger = &StandardLogger{baseLogger, name}
	standardLogger.Formatter = &logrus.JSONFormatter{}
  file, err := os.OpenFile("./logs/sample.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err == nil {
    standardLogger.SetOutput(file)
  }
	return standardLogger
}

// Declare variables to store log messages as new Events
var (
  invalidArgMessage      = Event{1, "Invalid arg: %s"}
  invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
  missingArgMessage      = Event{3, "Missing arg: %s"}
  withTracing      = Event{4, "Information : %s"}
)

func (l *StandardLogger) WithTracing(traceId string) {
  l.WithFields(logrus.Fields{
    "service_name": l.name,
    "trace_id": traceId,
  }).Info()
}

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
  l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
  l.WithFields(logrus.Fields{
    "service_name": l.name,
  }).Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
  l.Errorf(missingArgMessage.message, argumentName)
}