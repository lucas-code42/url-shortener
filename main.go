package main

import (
	"github.com/lucas-code42/url-shortner/api"
	"github.com/lucas-code42/url-shortner/configs"
)

func init() {
	configs.LoadConfigs()
}

func main() {
	api.InitializeServer()
}
