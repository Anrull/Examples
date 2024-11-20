package config

import (
    "gopkg.in/yaml.v2"
    "os"
)

type Config struct {
    ServerPort string `yaml:"server_port"`
    LogFile    string `yaml:"log_file"`
    Database   struct {
        Host     string `yaml:"host"`
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        DbName   string `yaml:"dbname"`
        Port     string `yaml:"port"`
        SSLMode  string `yaml:"sslmode"`
        Timezone string `yaml:"timezone"`
    } `yaml:"database"`
    Redis struct {
        Host     string `yaml:"host"`
        Password string `yaml:"password"`
        DB       int    `yaml:"db"`
    } `yaml:"redis"`
}

func LoadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }

    return &cfg, nil
}

