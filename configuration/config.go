package configuration

type Config struct {
	VersionString       string
	DebugMode           bool
	DataPath            string
	MinIoEndPoint       string
	MinIoAccessKey      string
	MinIoAcessKeySecret string
	MinIoSecure         bool
	MinIoBucket         string
}

var config Config

func Init(c Config) {
	config = c
}

func Get() Config {
	return config
}
