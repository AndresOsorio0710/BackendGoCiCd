package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PostgresConfig struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DBName   string `json:"DBName"`
	SSLMode  string `json:"SSLMode"`
}

type AppConfig struct {
	AppName  string         `json:"AppName"`
	Port     int            `json:"Port"`
	Postgres PostgresConfig `json:"Postgres"`
}

var Cfg AppConfig

func LoadConfig(basePath string) error {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "Development"
	}

	baseFile := filepath.Join(basePath, "appsettings.json")
	envFile := filepath.Join(basePath, fmt.Sprintf("appsettings.%s.json", env))

	if err := loadFromFile(baseFile, &Cfg); err != nil {
		return err
	}

	if err := loadFromFile(envFile, &Cfg); err != nil {
		fmt.Printf("No se encontró configuración para entorno %s (%s)\n", env, envFile)
	}

	overrideEnvVars()

	return nil
}

func loadFromFile(path string, cfg *AppConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(cfg)
}

func overrideEnvVars() {
	if host := os.Getenv("POSTGRES_HOST"); host != "" {
		Cfg.Postgres.Host = host
	}
	// Repetir para User, Password, etc.
}
