package proxy 

import (
    "fmt"
    "log"
    "os"

    "github.com/OpenVPN/openvpn"
)

func OpenVPNConnect() {

    // Load the openvpn configuration file
    config, err := openvpn.NewConfigFromFile("internal/config/openvpn.ovpn")
    if err != nil {
        log.Fatal(err)
    }

    // Create a new openvpn client 
    client, err := openvpn.NewClient(config, os.Stdout, os.Stderr)
    if err != nil {
        log.Fatal(err)
    }

    // Connect to the openvpn server
    err = client.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect()

    // Wait for the connection to be established
    fmt.Println("Connected to the VPN server")
    <-client.Wait()
}

