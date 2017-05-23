package app_fund

import (
	"ripple/app/handler"

	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
)

func App() {
	router := gin.Default()
	router.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "templates/",
		ContentType: "text/html; charset=utf-8",
	})

	router.GET("/taobao_fund/list", handler.Taobao)

	router.Run(":8010")
}
