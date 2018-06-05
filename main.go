package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mohan3d/proxy/proxy"
	"github.com/mohan3d/proxy/transform"
)

// URLParam constant
const URLParam = "url"

func setHeaders(headers http.Header, rw gin.ResponseWriter) {
	for k, vv := range headers {
		for _, v := range vv {
			rw.Header().Add(k, v)
		}
	}
}

func requestHandler(fetcher proxy.Fetcher) func(c *gin.Context) {
	return func(c *gin.Context) {
		url := c.Param(URLParam)[1:]

		// check for recursive request.

		// fetch the url.
		resp, err := fetcher.Fetch(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// read all content.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// transform content.
		content, err := transform.Transform(body, transform.HTML)
		if err != nil {
			panic(err)
		}

		// get allowed headers.
		headers := proxy.AllowedHeaders(resp.Header)

		// write response.
		c.Writer.WriteHeader(http.StatusOK)
		setHeaders(headers, c.Writer)
		c.Writer.Write(content)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	fetcher := proxy.URLFetcher{}

	router.GET("/*url", requestHandler(fetcher))
	router.Run(":" + port)
}
