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
