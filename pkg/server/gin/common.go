package gin

import "github.com/ziyoubiancheng/drivers/pkg/common"

type Cfg struct {
	Drivers struct {
		Server struct {
			Gin CallerCfg `toml:"gin"`
		} `toml:"server"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Mode            string
	Addr            string
	ReadTimeout     common.Duration
	WriteTimeout    common.Duration
	MaxHeaderBytes  int
	EnabledRecovery bool
	EnabledLogger   bool
	EnabledMetric   bool
}
