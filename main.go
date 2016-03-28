package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/VividCortex/godaemon"
	"github.com/gsora/slaidsciou/config"
	"github.com/gsora/slaidsciou/cycles"
)

var doThis func(imgs []os.FileInfo, conf config.SlaidsciouConfig)

func main() {
	// Software options
	foreground := flag.Bool("foreground", false, "don't fork in background, useful for debug purposes")
	flag.Parse()

	// Daemonize, or foreground :)
	if *foreground {
		godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	}

	godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	imgs, _ := ioutil.ReadDir(conf.WallpaperPath)

	if conf.Randomize {
		doThis = cycles.RandomCycle
	} else {
		doThis = cycles.NotRandomCycle
	}

	for {
		doThis(imgs, conf)
	}

}
