package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Router struct {
	Engine *gin.Engine
	DB     *gorm.DB
	Logger *logrus.Logger
}

func InitializeRouter(Engine *gin.Engine, DB *gorm.DB, Logger *logrus.Logger) *Router {
	r := &Router{
		Engine: Engine,
		DB:     DB,
		Logger: Logger,
	}
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
	return r
}

func (r *Router) SetupHandler() {
	r.Engine.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Server is still runing")
	})
}
