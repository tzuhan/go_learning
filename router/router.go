package router

import (
	"fmt"
	"time"

	_ "go_learning/docs"
	"go_learning/handler"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/unrolled/secure"
)

func SetupRouter(credentialType string) *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	if credentialType == "http" {
		v1 := router.Group("/api/v1")
		router.Use(cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:8080", "https://localhost:8080", "http://localhost:3002", "https://localhost:3002"},
			AllowedMethods:   []string{"GET"},
			AllowedHeaders:   []string{"Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
			AllowCredentials: true,
			Debug:            true,
			MaxAge:           12 * int(time.Hour),
		}))
		{
			v1.GET("/role", handler.GetRoles)
			v1.GET("/role/:id", handler.GetRole)
		}
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	} else if credentialType == "https" {
		secureMiddleware := secure.New(secure.Options{
			FrameDeny: true,
		})
		secureFunc := func() gin.HandlerFunc {
			return func(c *gin.Context) {
				err := secureMiddleware.Process(c.Writer, c.Request)

				// If there was an error, do not continue.
				if err != nil {
					c.Abort()
					return
				}

				// Avoid header rewrite if response is a redirection.
				if status := c.Writer.Status(); status > 300 && status < 399 {
					c.Abort()
				}
			}
		}()
		v1 := router.Group("/api/v1")
		router.Use(secureFunc)
		{
			v1.GET("/role", handler.GetRoles)
			v1.GET("/role/:id", handler.GetRole)
		}
		if mode := gin.Mode(); mode == gin.DebugMode {
			url := ginSwagger.URL(fmt.Sprintf("https://localhost:%d/swagger/doc.json", 8081))
			router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
		}
	}
	return router
}
