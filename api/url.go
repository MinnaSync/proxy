package api

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func ProxYURL(c *fiber.Ctx) error {
	urlParam := c.Params("*")
	parsedUrl, err := url.Parse(urlParam)
	if err != nil {
		println(err.Error())
		return nil
	}

	req, _ := http.NewRequest("GET", parsedUrl.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return nil
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return nil
	}

	c.Status(statusCode)
	c.Set("Content-Type", contentType)

	return c.Send(response)
}
