package main

import (
	"log"
	"os"

	"github.com/keekun/gompose/config"
	"github.com/keekun/gompose/proc"
)

func main() {
	conf, err := config.Load("")
	if err != nil {
		log.Fatal(err)
	}
	ps := proc.NewProcesses(conf, os.Stdout)

	ps.Spawn()
}
