package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

    api "lgame/internal/api"
)


func route(eng *gin.Engine) {
	group := eng.Group("/api")
	group.POST("/getgame", api.ApiGetGame)
	group.POST("/uploadresult", api.ApiUploadResult)
	group.POST("/top", api.ApiTop)
	group.POST("/login", api.ApiLogin)
	group.POST("/creategame", api.ApiAddGame)
	group.POST("/stopgame", api.ApiStopGame)
}

func main() {

	debug := flag.Bool("debug", false, "debug")

    port := flag.Int("port", 8080, "port")

    flag.Parse()

	if *debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	route(r)

	slog.Info("run", "port",  *port)
	err := r.Run(fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Error("gin Run", "error", err)
	}
}
