package webutil

import "github.com/gofiber/fiber/v2"

// MetaInfo is ...
type MetaInfo struct {
	Request Request `json:"request"`
	Data    any     `json:"data"`
}

// Request is ...
type Request struct {
	UserAgent string `json:"user_agent"`
	UserIP    string `json:"user_ip"`
}

// GetRequestInfo is ...
func GetRequestInfo(c *fiber.Ctx, data any) *MetaInfo {
	return &MetaInfo{
		Request: Request{
			UserAgent: string(c.Request().Header.UserAgent()),
			UserIP:    c.IP(),
		},
		Data: data,
	}
}
