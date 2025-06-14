package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

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

	bytes, err, _ := RequestGroup.Do(parsedUrl.String(), func() (any, error) {
		req, _ := http.NewRequest("GET", parsedUrl.String(), nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Log.Error("failed to request.", "error", err)
			return nil, err
		}
		defer resp.Body.Close()

		playlist, err := m3u8.Read(resp.Body)
		if err != nil {
			logger.Log.Error("failed to read playlist.", "error", err)
			return nil, err
		}

		// If the request is from local, assume it's http.
		var baseUrl string
		if c.IsFromLocal() {
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

		return buffer.Bytes(), nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(APIError{
			Message: "failed to proxy m3u8.",
			Error:   err.Error(),
		})
	}

	c.Set("Content-Type", "application/vnd.apple.mpegurl")
	return c.Send(bytes.([]byte))
}
