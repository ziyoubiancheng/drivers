package echosession

type Cfg struct {
	Drivers struct {
		Session struct {
			CallerCfg `toml:"echo"`
		} `toml:"session"`
	} `toml:"drivers"`
}

type CallerCfg struct {
	Name     string
	Size     int
	Network  string
	Addr     string
	Pwd      string
	Keypairs string
}
