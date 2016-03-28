package cycles

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gsora/slaidsciou/config"
)

// NotRandomCycle set pictures as background following the order gotten from ioutil.ReadDir
func NotRandomCycle(imgs []os.FileInfo, conf config.SlaidsciouConfig) {
	for _, img := range imgs {
		// gsettings set org.gnome.desktop.background picture-uri file://
		filename := "file://" + conf.WallpaperPath + "/" + img.Name()
		fmt.Println("Image: ", filename)
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", filename)
		cmd.Start()
		time.Sleep(time.Duration(conf.Delay) * time.Minute)
	}
}
