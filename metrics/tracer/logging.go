// Code generated by go generate; DO NOT EDIT.

package tracer

import (
	"github.com/stellarentropy/gravity-assist-common/logging"
)

var logger logging.Logger

func init() {
	logger = logging.Logger{Logger: logging.GetLogger().With().Str("component", componentName).Logger()}
}