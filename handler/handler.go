package handler

var options Options

type Options struct {
	VersionString string
}

func Init(o Options) {
	options = o
}

func CreateVersionString() string {
	return options.VersionString
}
