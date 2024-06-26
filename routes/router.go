package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"log"
	"task2/config"
	"task2/models/posts"
	"task2/routes/api/v1"
	"task2/routes/api/v2"
	"task2/routes/healthcheck"
	"task2/routes/middleware"
)

func SetupAppRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Logger(),
		middleware.ErrorHandler,
		gin.CustomRecovery(middleware.PanicHandler),
		middleware.CorsMiddleware(cfg))

	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("fatal SetTrustedProxies error: %v", err)
	}

	postsMemory := posts.FromFileToMemoryDs(cfg.InitialDataFile)
	v1.NewPostsController(router.Group("/api/v1/posts"), postsMemory)

	humaConfig := huma.DefaultConfig("Blog API", "0.2.0")
	humaConfig.Servers = []*huma.Server{{URL: cfg.UrlOrigin}}
	api := humagin.New(router, humaConfig)
	v2.NewPostsController(api, "/api/v2/posts", postsMemory)

	router.GET("/healthcheck", healthcheck.PingHandler)

	return router
}

func RunApp(cfg *config.Config) {
	if err := SetupAppRouter(cfg).Run(cfg.Addr); err != nil {
		log.Fatalf("fatal Engine.Run error: %v", err)
	}
}
