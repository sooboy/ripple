package handler

import (
	"fmt"
	db "ripple/database"
	"ripple/fund"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Taobao(c *gin.Context) {
	var results []fund.Result
	err := db.Session.DB("test").C("ripple.fund").Find(nil).All(&results)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, results)
}
