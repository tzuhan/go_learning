package main

import (
	"github.com/tzuhan/go_learning/router"
	//"github.com/unrolled/secure" //HTTPS
)

func main() {

	r := router.SetupRouter()
	//router.Use(TlsHandler())
	r.Run(":8080")
	//router.RunTLS(":8080", "ssl.pem", "ssl.key")
}

/*
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:3002",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}*/
