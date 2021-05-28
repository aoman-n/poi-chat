package redis

import (
	"fmt"

	"github.com/laster18/poi/api/src/config"
)

const (
	EventSet     = "set"
	EventExpired = "expired"
	EventDel     = "del"
)

var (
	KeySpace = fmt.Sprintf("__keyspace@%d__", config.Conf.Redis.Db)
	KeyEvent = fmt.Sprintf("__keyevent@%d__", config.Conf.Redis.Db)
)
