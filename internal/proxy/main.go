package proxy

import (
	"fmt"

	"github.com/mmalcherczyk/go-proxy"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.All("/*", proxy.ProxyHandler)

	fmt.Println("Starting proxy server on http://localhost:3000")
	app.Listen(":3000")
}
