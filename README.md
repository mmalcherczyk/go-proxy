PROXY-GO server project 
=======================
This project is a simple server that can be used to proxy requests to other servers.

File and folder structure
--------------------------
```proxy-server/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── proxy/
│   │   └── proxy.go
│   └── vpn/
│       └── vpn.go
├── pkg/
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── go.mod
└── go.sum

```
1. cmd/main.go: The main entry point of the application.
2. internal/config/config.go: The configuration file for the application.
3. internal/proxy/proxy.go: The proxy server implementation.
4. internal/vpn/vpn.go: The VPN server implementation.
5. pkg/logger/logger.go: The logger implementation.
6. pkg/utils/utils.go: The utility functions.
7. go.mod: The module file for the application.
8. go.sum: The module checksum file for the application.
