package main

import (
	"flag"
	"github.com/maocatooo/toss/oss"
)

var configPath = flag.String("config", "config.yaml", "")

func main() {
	flag.Parse()
	conf := oss.ConfigFromFile(*configPath)
	oss.Upload(conf)
}
