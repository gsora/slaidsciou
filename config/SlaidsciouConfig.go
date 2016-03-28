package config

// SlaidsciouConfig represent a JSON configuration file
type SlaidsciouConfig struct {
	WallpaperPath string `json:"wallpaper-path"`
	Randomize     bool   `json:"randomize"`
	Delay         int    `json:"delay"`
}
