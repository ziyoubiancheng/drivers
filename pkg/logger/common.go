package logger

type Cfg struct {
	Drivers struct {
		Logger map[string]CallerCfg `toml:"logger"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Debug bool
	Level string
	Path  string
}
