package routers

import (
	"time"

	"github.com/Reshnyak/AcornsTask/internal/configs"
	"github.com/Reshnyak/AcornsTask/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func InitRouter(cfg *configs.Config) *gin.Engine {
	initLimiter(cfg.HTTPServer.RpcLimit)
	router := gin.New()
	//Тонна миддлварек :)
	router.Use(gin.Recovery())
	// router.Use(otelgin.Middleware("alise-todo"))
	router.Use(gin.Logger())
	router.Use(rateLimit())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  cfg.Cors.AllowAllOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		ExposeHeaders:    cfg.Cors.ExposeHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
		MaxAge:           cfg.Cors.MaxAge * time.Hour,
	}))

	router.POST("/input", handlers.AlicePost())
	// Подключение Swagger yaml
	router.StaticFile("/docs", "./docs/swagger.yaml")
	return router
}
