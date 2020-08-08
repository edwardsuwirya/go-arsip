package main

import (
	"config"
	"deliveries"
)

type app struct {
	config *config.Config
}

func (a app) run() {
	if a.config.ConsoleGui == 1 {
		deliveries.InitUi(a.config).Run()
	}
}

func newApp(config *config.Config) *app {
	return &app{
		config,
	}
}
func main() {
	//go run main <{en|id}> <file_name.{json|csv}> <{1|0}>
	c := config.NewConfig()
	newApp(c).run()
}
