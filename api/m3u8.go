package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/MinnaSync/proxy/internal/logger"
	"github.com/etherlabsio/go-m3u8/m3u8"
	"github.com/gofiber/fiber/v2"
)

func ProxYM3U8(c *fiber.Ctx) error {
	urlParam := c.Params("*")
	parsedUrl, err := url.Parse(urlParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(APIError{
			Message: "an invalid url was provided.",
			Error:   err.Error(),
		})
	}

	req, _ := http.NewRequest("GET", parsedUrl.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Error("failed to request.", "error", err)

		return c.Status(fiber.StatusInternalServerError).JSON(APIError{
			Message: "failed to request.",
			Error:   err.Error(),
		})
	}
	defer resp.Body.Close()

	playlist, err := m3u8.Read(resp.Body)
	if err != nil {
		logger.Log.Error("failed to read playlist.", "error", err)

		return c.Status(fiber.StatusInternalServerError).JSON(APIError{
			Message: "failed to read playlist.",
			Error:   err.Error(),
		})
	}

	// If the environment is development, assume the protocol is http.
	var baseUrl string
	if os.Getenv("ENVIRONMENT") == "development" {
		baseUrl = fmt.Sprintf("http://%s", c.Hostname())
	} else {
		baseUrl = fmt.Sprintf("https://%s", c.Hostname())
	}

	for _, item := range playlist.Items {
		switch item := item.(type) {
		case *m3u8.KeyItem:
			proxyURI := fmt.Sprintf("%s/url/%s", baseUrl, *item.Encryptable.URI)
			item.Encryptable.URI = &proxyURI
		case *m3u8.PlaylistItem:
			proxyURI := fmt.Sprintf("%s/url/%s", baseUrl, item.URI)
			item.URI = proxyURI
		case *m3u8.SegmentItem:
			proxyURI := fmt.Sprintf("%s/url/%s", baseUrl, item.Segment)
			item.Segment = proxyURI
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString(playlist.String())

	c.Set("Content-Type", "application/vnd.apple.mpegurl")
	return c.Send(buffer.Bytes())
}
