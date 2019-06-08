package config

// application : struct to hold application level configs
type application struct {
	Name       string `toml:"app_name"`
	Mode       string `toml:"app_mode"`
	ListenPort int    `toml:"listen_port"`
	ListenIP   string `toml:"listen_ip"`
	LogPath    string `toml:"log_path"`
}
