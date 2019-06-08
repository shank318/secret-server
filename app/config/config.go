// Package config will manage all application level configurations
// config file will be taken based on the application environment
// all the configuration available in router file will be overwritten
// this will be immutable as it always provides the value of the struct
package config

const (
	// FilePath - relative path to the config directory
	FilePath = "%s/conf/%s"

	// DefaultFilename - Filename format of default config file
	DefaultFilename = "env.default.toml"

	// EnvFilename - Filename format of env specific config file
	EnvFilename = "env.%s.toml"
)

var (
	// config : this will hold all the application configuration
	config AppConfig
)

// appConfig global configuration struct definition
type AppConfig struct {
	Application application `toml:"application`
	NewRelic    newrelic    `toml:"newrelic"`
	AuthUser    authUser    `toml:"authuser"`
	Prometheus  prometheus  `toml:"prometheus"`
}

// LoadConfig will load the configuration available in the cnf directory available in basePath
// conf file will takes based on the env provided
func LoadConfig(basePath string, env string) {
	// reading conf based on default environment
	loadConfigFromFile(basePath, DefaultFilename, "")

	// reading env file and override conf values; if env file exists
	loadConfigFromFile(basePath, EnvFilename, env)
}

// GetConfig : will give the struct as value so that the actual conf doesn't get tampered
func GetConfig() AppConfig {
	return config
}
