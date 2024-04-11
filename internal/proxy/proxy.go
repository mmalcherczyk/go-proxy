package proxy

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v3"
)

func ProxyHandler(c *fiber.Ctx) error {
	// Get the URL from the request
	requestedURL := c.Request().URL.String()

	// Parse the URL
	url, err := url.Parse(requestedURL)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid URL")
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest(c.Request().Method, url.String(), c.Request().Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error creating request")
	}

    // Copy the headers from the original request
    for key, values := range c.Request().Header {
        for _, value := range values {
            req.Header.Add(key, value)
        }
    }
    // Send the request to the server
    resp, err := client.Do(req)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Error sending request")
    }
    defer resp.Body.Close()

    // Copy the headers from the response
    for key, values := range resp.Header {
        for _, value := range values {
            c.Response().Header.Add(key, value)
        }
    }

    // Copy the response body to the original response
    _, err = io.Copy(c.Response().Body, resp.Body)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Error copying response")
    }

    return nil

}






}
