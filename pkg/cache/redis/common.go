package redis

import (
	"github.com/ziyoubiancheng/drivers/pkg/common"
)

type Cfg struct {
	Drivers struct {
		Redis map[string]CallerCfg `toml:"redis"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Debug bool

	Network        string // tcp
	Addr           string // 127.0.0.1:6379
	DB             int
	Password       string
	ConnectTimeout common.Duration
	ReadTimeout    common.Duration
	WriteTimeout   common.Duration

	MaxIdle     int
	MaxActive   int
	IdleTimeout common.Duration
	Wait        bool
}
