package stat

import "github.com/ziyoubiancheng/drivers/pkg/common"

type Cfg struct {
	Drivers struct {
		Server struct {
			Stat CallerCfg `toml:"stat"`
		} `toml:"server"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Addr         string
	ReadTimeout  common.Duration
	WriteTimeout common.Duration
}
