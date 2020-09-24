package config

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config is a thing
type Config struct {
	RedisAddr string `yaml:"redis_addr"`
	RedisPass string `envconfig:"REDIS_PASS"`
	// Commands       map[string]*Command        `json:"commands"`
	// Triggers       map[string]Trigger         `json:"triggers"`
	// EnabledModules []string `json:"enabledModules"`
	// DatabasePath   string   `json:"databasePath"`
	// WebPath   string `json:"webPath"`
	// MediaPath string `json:"mediaPath"`
	// ModuleConfig   map[string]json.RawMessage `json:"moduleConfig"`
}

// LoadConfig loads the config
func LoadConfig(rdr io.Reader) (*Config, error) {
	bts, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}
	cfg := new(Config)
	if err := yaml.Unmarshal(bts, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
	// dec := json.NewDecoder(r)
	// if err := dec.Decode(&config); err != nil {
	// 	return err
	// }

	// for key := range config.Commands {
	// 	cmd := config.Commands[key]
	// 	cmd.Name = key
	// }

	// return nil
}
