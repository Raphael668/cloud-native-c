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

func Withdrawal(ctx *gin.Context) {
	response := ResBody{}
	response.ResCode = RES_OK

	ctx.JSON(http.StatusOK, response)
}
