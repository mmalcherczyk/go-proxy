package proxy

import (
    "fmt"
    "log"
    "os" 
)

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
