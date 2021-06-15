package alog

import (
	"fmt"
	"io"
	"os"

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

var DefaultLogger = New()

type Log struct {
	zerolog zerolog.Logger
}

func New(options ...Option) *Log {
	cfg := newDefaultConfig()

	for _, option := range options {
		option.Apply(cfg)
	}

	var output io.Writer
	output = os.Stdout
	if !cfg.WithJSON {
		output = zerolog.ConsoleWriter{Out: os.Stdout}
	}

	zerolog.ErrorStackMarshaler = marshalStack
	lc := zerolog.New(output).With().Timestamp()

	if cfg.RequestID != "" {
		lc = lc.Str("requestId", cfg.RequestID)
	}

	if cfg.IP != "" {
		lc = lc.Str("ip", cfg.IP)
	}

	if cfg.User != nil {
		uname := cfg.User.Name
		if uname == "" {
			uname = "unknown"
		}
		lc = lc.Str("userName", uname)

		uid := cfg.User.ID
		if uid == "" {
			uid = "unknown"
		}
		lc = lc.Str("userId", uid)
	}

	if cfg.WithCaller {
		lc = lc.CallerWithSkipFrameCount(3)
	}

	logger := lc.Logger().Level(level(cfg.Level))

	return &Log{logger}
}

// marshalStack エラースタックトレースログ出力をカスタマイズするための関数
//
// 下記パラメータを出力するようにカスタマイズしている
// msg: エラーメッセージ
// func: エラーが発生した関数名
// line: エラーが発生した行番号
//
// 出力例:
// "stack": [
// 	  {
// 		"message": "errorです！",
// 		"func": "github.com/laster18/poi/api/src/delivery/graphql.(*queryResolver).Rooms",
// 		"line": "/app/src/delivery/graphql/room.resolver.go:52"
//    },
//    ...
// ]
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
