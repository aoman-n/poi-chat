package alog

import (
	"github.com/laster18/poi/api/src/config"
	"github.com/rs/zerolog"
)

type LogLevel string

const (
	LevelInfo  LogLevel = "info"
	LevelDebug LogLevel = "debug"
)

func level(l LogLevel) zerolog.Level {
	switch l {
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelDebug:
		return zerolog.DebugLevel
	default:
		return zerolog.InfoLevel
	}
}

type Config struct {
	// RequestID リクエストの識別子ID。空文字列の場合は出力しない
	RequestID string
	// IP リクエスト元のIP。空文字列の場合は出力しない
	IP string
	// IP リクエストユーザー情報。nilの場合は出力しない
	User *User
	// WithJSON trueにするとJSONでログを出力
	WithJSON bool
	// WithCaller trueにするとログ出力したファイルと行数を表示
	WithCaller bool
	// Level ログ出力レベルの設定。　LevelInfo or LevelDebug デフォルトの設定は LevelInfo
	Level LogLevel
}

func newDefaultConfig() *Config {
	return &Config{
		RequestID:  "",
		IP:         "",
		User:       nil,
		WithJSON:   config.IsProd(),
		WithCaller: config.IsDev(),
		Level:      LogLevel(config.Conf.LogLevel),
	}
}

type Option interface {
	Apply(*Config)
}

type withRequetID string

func (w withRequetID) Apply(cfg *Config) {
	cfg.RequestID = string(w)
}

func WithRequetID(id string) Option {
	return withRequetID(id)
}

type withIP string

func (w withIP) Apply(cfg *Config) {
	cfg.IP = string(w)
}

func WithIP(ip string) Option {
	return withIP(ip)
}

type User struct {
	ID   string
	Name string
}

func (u *User) Apply(cfg *Config) {
	cfg.User = u
}
