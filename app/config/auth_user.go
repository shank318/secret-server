package config

import "time"

// authUser: All the users that are going to interact with Scrooge
type authUser struct {
	API AuthUserConfig `toml:"api"`
}

// AuthUserConfig: Stores credentials config of a given authUser
type AuthUserConfig struct {
	UserName string `toml:"username"`
	Password string `toml:"password"`
}

type Auth struct {
	Key    string `env:"key"`
	Secret string `env:"secret"`
}

type Endpoint struct {
	URL     string            `toml:"url"`
	Method  string            `toml:"method"`
	Timeout time.Duration     `toml:"timeout"`
	Auth    Auth              `toml:"auth"`
	Headers map[string]string `toml:"headers"`
}
