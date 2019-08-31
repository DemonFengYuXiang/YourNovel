package http

import (
	"github.com/gin-gonic/gin"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/db/redis"
	"yournovel/yournovel/middleware"
	"yournovel/yournovel/tool"
)

func Init() {
	g := gin.Default()
	initRouter(g)
	err := g.Run(tool.GetHost())
	if err != nil {
		panic(err)
	}
}

// 初始化路由
func initRouter(engine *gin.Engine) {
	conf.InitConfig()
	redis.InitRedisClient()
	initConfig(engine)
	initGetRouter(engine)
	initPostRouter(engine)
}

func initConfig(engine *gin.Engine) {
	engine.LoadHTMLGlob("./yournovel/view/**/**/*")
	engine.Static("/assets", "./yournovel/static")
}

// 初始化Get路由
func initGetRouter(engine *gin.Engine) {
	engine.GET("/", middleware.RequestMiddlewareWrapper(home, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))

	engine.GET("/chapter", middleware.RequestMiddlewareWrapper(novelChapter, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))

	engine.GET("/search", middleware.RequestMiddlewareWrapper(novelSearch, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))

	engine.GET("/content", middleware.RequestMiddlewareWrapper(novelContent, middleware.MyMiddlewareOption{
		IsAuth: false,
	}))
}

// 初始化Post路由
func initPostRouter(engine *gin.Engine) {

}
