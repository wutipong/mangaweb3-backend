package handler

import "github.com/wutipong/mangaweb3-backend/ent"

var options Options

type Options struct {
	VersionString string
	EntClient     *ent.Client
}

func Init(o Options) {
	options = o
}

func CreateVersionString() string {
	return options.VersionString
}

func EntClient() *ent.Client {
	return options.EntClient
}
