package app

type Cfg struct {
	Drivers struct {
		App CallerCfg `toml:"app"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Name    string
	Version string
	Env     string
}
