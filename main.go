package main

import (
	db "ripple/database"
	"ripple/fund"
	"ripple/utils"
)

func init() {
	db.Init()
}
func main() {
	exites := make(chan bool)
	// 定期更新资金
	utils.NewTask(2*60*60, fund.TaobaoFundIn)

	// time.Sleep(60 * time.Second)
	<-exites
}
