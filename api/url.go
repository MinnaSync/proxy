package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/MinnaSync/proxy/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func ProxYURL(c *fiber.Ctx) error {
	logger.Log.Debug(fmt.Sprintf("Request Headers: \n%s", &c.Request().Header))

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

	statusCode := resp.StatusCode
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Log.Error("failed to request.", "error", err)

		return c.Status(fiber.StatusInternalServerError).JSON(APIError{
			Message: "failed to read response.",
			Error:   err.Error(),
		})
	}

	c.Status(statusCode)
	c.Set("Content-Type", contentType)

	return c.Send(response)
}
