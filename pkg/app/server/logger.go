package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type APILog struct {
	Version      string
	RequestURI   string
	Method       string
	Duration     time.Duration
	DurationText string
	InfoTxt      string
	ErrorTxt     string
	Error        error
}

func Logger(ctx *gin.Context) {
	now := time.Now()

	ctx.Next()

	log := APILog{}

	log.RequestURI = ctx.Request.RequestURI
	log.Method = ctx.Request.Method
	log.Duration = time.Since(now)
	log.DurationText = fmt.Sprintf("%v", log.Duration)

	fmt.Println(log)

	// if theError != nil {
	// 	log.Error = theError
	// 	log.ErrorTxt = log.Error.Error()
	// 	Warning(log)
	// } else {
	// 	Info(log)
	// }

}
