package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/etherlabsio/go-m3u8/m3u8"
	"github.com/gofiber/fiber/v2"
)

func ProxYM3U8(c *fiber.Ctx) error {
	urlParam := c.Params("*")
	parsedUrl, err := url.Parse(urlParam)
	if err != nil {
		return nil
	}

	req, _ := http.NewRequest("GET", parsedUrl.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	playlist, err := m3u8.Read(resp.Body)
	if err != nil {
		return nil
	}

	protocol := c.Get("x-forwarded-proto")
	if protocol == "" {
		protocol = "http"
	}

	baseUrl := fmt.Sprintf("%s://%s", protocol, c.Hostname())
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
