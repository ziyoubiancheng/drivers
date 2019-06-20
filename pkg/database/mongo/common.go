package mongo

type Cfg struct {
	Drivers struct {
		Mongo map[string]CallerCfg `toml:"mongo"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Debug bool

	URL      string
	Database string
}
