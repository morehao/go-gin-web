package context

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetReqBody(c *gin.Context) (string, error) {
	if c.Request.Body == nil {
		return "", nil
	}
	byteBody, err := c.GetRawData()
	if err != nil {
		return "", err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))
	reqBody := string(byteBody)
	return reqBody, nil
}

func GetReqQuery(c *gin.Context) (reqQuery string) {
	reqQuery = c.Request.URL.RawQuery
	return reqQuery
}

func GetCookie(c *gin.Context) string {
	cookie := ""
	for _, c := range c.Request.Cookies() {
		cookie += fmt.Sprintf("%s=%s&", c.Name, c.Value)
	}
	return strings.TrimRight(cookie, "&")
}
func GetHeader(c *gin.Context) string {
	header := ""
	for k, v := range c.Request.Header {
		header += fmt.Sprintf("%s=%s&", k, v)
	}
	return strings.TrimRight(header, "&")
}
