package config

// newrelic: struct to hold newrelic config
type newrelic struct {
	Enabled      bool   `toml:"enabled"`
	LicenceKey   string `toml:"licence_key"`
	AppName      string `toml:"app_name"`
	HighSecurity bool   `toml:"high_security"`
}
