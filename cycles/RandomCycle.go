package cycles

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/gsora/slaidsciou/config"
)

// RandomCycle cycles through random photos and sets them as wallpaper
func RandomCycle(imgs []os.FileInfo, conf config.SlaidsciouConfig) {
	imgsSize := len(imgs)
	for i := 0; i < imgsSize; i++ {
		rand.Seed(time.Now().Unix())
		filename := "file://" + conf.WallpaperPath + "/" + imgs[rand.Intn(imgsSize)].Name()
		fmt.Println("Image: ", filename)
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", filename)
		cmd.Start()
		time.Sleep(time.Duration(conf.Delay) * time.Minute)
	}
}
