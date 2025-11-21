package userconfig

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

type Config struct {
	// ServerURL string `json:"serverUrl"`
	RetroachievmentKey string
	AutoValidate       bool
}

func DefaultConfig() Config {
	return Config{
		// ServerURL: "https://api.example.com",
	}
}

type ConfigStore struct {
	configPath string
}

func NewConfigStore() (*ConfigStore, error) {
	configFilePath, err := xdg.ConfigFile("Romscraper/config.json")
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	return &ConfigStore{
		configPath: configFilePath,
	}, nil
}

func (s *ConfigStore) Get() (Config, error) {
	_, err := os.Stat(s.configPath)
	if os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	dir, fileName := filepath.Split(s.configPath)
	if len(dir) == 0 {
		dir = "."
	}

	buf, err := fs.ReadFile(os.DirFS(dir), fileName)
	if err != nil {
		return Config{}, fmt.Errorf("could not read the configuration file: %w", err)
	}

	if len(buf) == 0 {
		return DefaultConfig(), nil
	}

	cfg := Config{}
	if err := json.Unmarshal(buf, &cfg); err != nil {
		return Config{}, fmt.Errorf("configuration file does not have a valid format: %w", err)
	}

	return cfg, nil

}

func (s *ConfigStore) Save(cfg Config) error {
	jsoncfg, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("Not valid Config: %w", err)
	}

	err = os.WriteFile(s.configPath, jsoncfg, 0644)
	if err != nil {
		return fmt.Errorf("Couldn't write a file: %w", err)
	}

	return nil
}

func main() {
	store, err := NewConfigStore()
	if err != nil {
		fmt.Printf("could not initialize the config store: %v\n", err)
		return
	}

	fmt.Println(store.configPath)

	cfg, err := store.Get()
	if err != nil {
		fmt.Printf("could not retrieve the configuration: %v\n", err)
		return
	}
	fmt.Printf("config: %v\n", cfg)

	newcfg := Config{RetroachievmentKey: "asdasd", AutoValidate: false}

	err = store.Save(newcfg)

	if err != nil {
		fmt.Printf("couldn't save configuration: %v\n", err)
	}

	err = store.Save(newcfg)
}
