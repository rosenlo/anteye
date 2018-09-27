package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RosenLo/anteye/g"
	"github.com/RosenLo/anteye/http"
	"github.com/RosenLo/anteye/monitor"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)

	// monitor
	monitor.Start()

	// http
	http.Start()

	select {}
}
