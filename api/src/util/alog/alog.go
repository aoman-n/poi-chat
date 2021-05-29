package alog

import (
	"fmt"
	"io"
	"os"

	"github.com/laster18/poi/api/src/config"
	"github.com/laster18/poi/api/src/util/aerrors"
	"github.com/rs/zerolog"
)

type Logger interface {
	Warn(msg string)
	Warnf(format string, v ...interface{})
	WarnWithErr(err error, msg string)
	WarnfWithErr(err error, format string, v ...interface{})
	Info(msg string)
	Infof(format string, v ...interface{})
	Debug(msg string)
	Debugf(format string, v ...interface{})
}

var DefaultLogger Logger

func init() {
	DefaultLogger = Default()
}

type Log struct {
	zerolog zerolog.Logger
}

type User struct {
	ID   string
	Name string
}

type Conf struct {
	RequestID string
	User      *User
	IP        string
	IsJSON    bool
	IsCaller  bool
	// Level: "info" or "debug", default is "info"
	Level string
}

func create(c *Conf) *Log {
	var output io.Writer
	output = os.Stdout
	if !c.IsJSON {
		output = zerolog.ConsoleWriter{Out: os.Stdout}
	}
	zerolog.ErrorStackMarshaler = marshalStack

	lc := zerolog.New(output).With().Timestamp()
	lc = lc.Str("requestId", c.RequestID)

	if c.User != nil {
		uname := c.User.Name
		if uname == "" {
			uname = "unknown"
		}
		lc = lc.Str("userName", uname)

		uid := c.User.ID
		if uid == "" {
			uid = "unknown"
		}
		lc = lc.Str("userId", uid)
	}

	if c.IP != "" {
		lc = lc.Str("ip", c.IP)
	}

	if c.IsCaller {
		lc = lc.CallerWithSkipFrameCount(3)
	}

	logger := lc.Logger().Level(level(c.Level))

	return &Log{logger}
}

func New(c *Conf) *Log {
	return create(c)
}

func Default() *Log {
	return create(&Conf{
		IsJSON:    !config.IsDev(),
		RequestID: "Unknown",
		User:      nil,
		IsCaller:  false,
		Level:     config.Conf.LogLevel,
	})
}

func level(l string) zerolog.Level {
	switch l {
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}

func marshalStack(err error) interface{} {
	aerr, ok := err.(*aerrors.ErrApp)
	if !ok {
		return nil
	}

	out := []map[string]string{}
	for _, s := range aerr.StackTrace() {
		o := map[string]string{}
		if s.Msg != "" {
			o["message"] = s.Msg
		}
		if s.Func != "" {
			o["func"] = s.Func
		}
		if s.File != "" {
			o["line"] = fmt.Sprintf("%s:%d", s.File, s.Line)
		}

		out = append(out, o)
	}

	return out
}

func (l *Log) Warn(msg string) {
	l.zerolog.Warn().Msg(msg)
}

func (l *Log) Warnf(format string, v ...interface{}) {
	l.zerolog.Warn().Msgf(format, v...)
}

func (l *Log) WarnWithErr(e error, msg string) {
	l.zerolog.Warn().Stack().Err(e).Msg(msg)
}

func (l *Log) WarnfWithErr(e error, format string, v ...interface{}) {
	l.zerolog.Warn().Stack().Err(e).Msgf(format, v...)
}

func (l *Log) Info(msg string) {
	l.zerolog.Info().Msg(msg)
}

func (l *Log) Infof(format string, v ...interface{}) {
	l.zerolog.Info().Msgf(format, v...)
}

func (l *Log) Debug(msg string) {
	l.zerolog.Debug().Msg(msg)
}

func (l *Log) Debugf(format string, v ...interface{}) {
	l.zerolog.Debug().Msgf(format, v...)
}
