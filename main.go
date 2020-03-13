package main

import (
	"github.com/tzuhan/go_learning/router"
	//"github.com/unrolled/secure" //HTTPS
)

// @title GetRole API
// @version 1.0
// @description This is a get role server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email stoola08@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
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
