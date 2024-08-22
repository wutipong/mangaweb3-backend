package configuration

type Config struct {
	VersionString string
	DebugMode     bool
	DataPath      string
}

var config Config

func Init(c Config) {
	config = c
}

func Get() Config {
	return config
}
