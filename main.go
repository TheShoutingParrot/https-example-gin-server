// An example https server with gin
// Inspired by the following tutorial: 
// https://medium.com/it-wonders-web/build-a-locally-trusted-https-server-with-golang-mkcert-in-20-mins-4d7ef68f5dac

package main

import (
	"log"
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func main () {
	var certPath, keyPath string

	// Parse arguments
	flag.StringVar(&certPath, "crt", "certs/localhost.pem", "the path of the SSL cert")
	flag.StringVar(&keyPath, "key", "certs/localhost-key.pem", "the path of the SSL key")
	flag.Parse()

	log.Println("path of the ssl cert", certPath)
	log.Println("path of the ssl key", keyPath)

	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/", ping)

	log.Fatal(r.RunTLS("localhost:8080", certPath, keyPath))
}
