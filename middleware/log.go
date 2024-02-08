package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type logger struct{}

var Logger logger

func structToJSON(data any) (string, error) {
	result, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (*logger) CustomLogger() echo.MiddlewareFunc {
	return echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			return strings.Contains(path, "/styles/") || strings.Contains(path, "/scripts/")
		},
		LogRequestID: true,
		LogProtocol:  true,
		LogMethod:    true,
		LogStatus:    true,
		LogHost:      true,
		LogUserAgent: true,
		LogRoutePath: true,
		LogURI:       true,
		LogURIPath:   true,
		LogLatency:   true,
		LogError:     true,
		LogReferer:   true,
		LogHeaders: []string{
			http.CanonicalHeaderKey("X-Request-Id"),
			http.CanonicalHeaderKey("X-Forwarded-For"),
			http.CanonicalHeaderKey("X-Real-Ip"),
			http.CanonicalHeaderKey("Accept"),
			http.CanonicalHeaderKey("Content-Type"),
			http.CanonicalHeaderKey("Accept-Encoding"),
		},
		LogValuesFunc: func(c echo.Context, values echomiddleware.RequestLoggerValues) error {
			type logValues struct {
				RequestID string
				Protocol  string
				Method    string
				Status    int
				Host      string
				UserAgent string
				RoutePath string
				URI       string
				URIPath   string
				Latency   string
				Error     error
				Referer   string
				Headers   map[string][]string
			}

			result, err := structToJSON(&logValues{
				RequestID: values.RequestID,
				Protocol:  values.Protocol,
				Method:    values.Method,
				Status:    values.Status,
				Host:      values.Host,
				UserAgent: values.UserAgent,
				RoutePath: values.RoutePath,
				URI:       values.URI,
				URIPath:   values.URIPath,
				Latency:   values.Latency.String(),
				Error:     values.Error,
				Referer:   values.Referer,
				Headers:   values.Headers,
			})

			if err != nil {
				return err
			}

			log.Println(result)
			return nil
		},
	})
}
