package config

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/user"
)

// WARNING: never directly use this variable before calling getUserHome()
var configFilePath = "/.config/slaidsciou.conf"

func getUserHome() {
	cU, _ := user.Current()
	configFilePath = cU.HomeDir + configFilePath
}

// LoadConfig loads a configuration file from the standard path, defined by "configFilePath"
func LoadConfig() (SlaidsciouConfig, error) {
	getUserHome()

	f, err := os.Open(configFilePath)
	defer f.Close()

	if err != nil {
		return SlaidsciouConfig{}, err
	}

	fReader := bufio.NewReader(f)

	buf := new(bytes.Buffer)
	buf.ReadFrom(fReader)

	var conf SlaidsciouConfig
	json.Unmarshal(buf.Bytes(), &conf)

	//
	// Configuration file sanity checks
	//

	// Empty configuration file?
	if conf == (SlaidsciouConfig{}) {
		e := errors.New("malformed configuration file")
		return SlaidsciouConfig{}, e
	}

	// Empty wallpaper path?
	if conf.WallpaperPath == "" {
		e := errors.New("cannot continue without wallpaper path")
		return SlaidsciouConfig{}, e
	}

	// Empty delay?
	if conf.Delay == 0 {
		e := errors.New("cannot continue without a wallpaper change delay")
		return SlaidsciouConfig{}, e
	}

	// Return everything!
	return conf, nil
}
