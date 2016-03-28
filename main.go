package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gsora/slaidsciou/config"
	"github.com/gsora/slaidsciou/cycles"
)

var doThis func(imgs []os.FileInfo, conf config.SlaidsciouConfig)

func main() {
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
