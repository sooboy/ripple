package main

import (
	"ripple/db"
	"ripple/fund"
)

func init() {
	db.Init()
}
func main() {
	fund.TaobaoFundIn()
}
