package routers

import (
	"time"

	"github.com/Reshnyak/AcornsTask/internal/loggers"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
)

func initLimiter(rpc int) {
	limit = ratelimit.New(rpc)
}
func rateLimit() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		loggers.Log().Warn().Msg(now.Sub(prev).String())
		prev = now
	}
}
