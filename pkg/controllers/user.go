package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deposit(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

func Withdraw(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}

//users/0/deposit
func DepositTo(ctx *gin.Context) { //TODO:

}

func WithdrawFrom(ctx *gin.Context) { //TODO:

}
