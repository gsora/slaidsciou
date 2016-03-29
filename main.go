package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/VividCortex/godaemon"
	"github.com/gsora/slaidsciou/config"
	"github.com/gsora/slaidsciou/cycles"
)

var doThis func(imgs []os.FileInfo, conf config.SlaidsciouConfig)

func main() {
	// Software options
	foreground := flag.Bool("foreground", false, "don't fork in background, useful for debug purposes")
	flag.Parse()

	// First things first: check for any other living instances of yourself.
	// If any, exit.
	cmd := exec.Command("pgrep", "slaidsciou")

	myPid := strconv.Itoa(os.Getpid())
	out, _ := cmd.CombinedOutput()
	k := strings.Split(string(out), "\n")

	if myPid != k[0] {
		fmt.Println("Slaidsciou already running, exiting...")
		return
	}

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
