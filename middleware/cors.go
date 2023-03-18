package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors2() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               // request method
		origin := c.Request.Header.Get("Origin") // request header

		var headerKeys []string // declaration request header keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                        // This is to allow access to all domains
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE") // all methods the server supports cross-domain request, in order to avoid multiple requests Viewed times' pre subject 'request
			// header type
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// allow cross-domain settings return to the other sub-segments
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, Expires, Last -Modified, Pragma, FooBar ") // key settings allow cross-domain browser may resolve
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                                    // cache request information in seconds
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                           // cookie with the need for cross-domain requests information defaults to true
			c.Set("content-type", "application / json")                                                                                                                                                                     // set the response format is json
		}

		// release all the OPTIONS method
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// process the request
		c.Next() // processing request
	}
}
