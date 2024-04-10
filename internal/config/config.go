package config

import (
    "fmt"
    "os"
    "gopkg.in/yaml.v2"
)

type VPNConfig struct {
    ServerAddr  string `yaml:"server_addr"`
    Username    string `yaml:"username"`
    Password    string `yaml:"password"`
}

func LoadConfig() (*VPNConfig, error) {
    // Odczytaj konfiguracje z pliku oraz zwróć obiekt VPNConfig
    configFile, error := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    
    defer configFile.Close()
    
    var config VPNConfig
    err := json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
