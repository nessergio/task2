package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"task2/config"
)

func CorsMiddleware(cfg *config.Config) gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{cfg.UrlOrigin}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	return cors.New(corsConfig)
}
