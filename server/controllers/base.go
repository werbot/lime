package controllers

import (
	"github.com/gin-gonic/gin"
)

type requestLicense struct {
	StripeID string `json:"customer"`
	License  string `json:"license"`
}

type ResponseData struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func respondJSON(g *gin.Context, status int, msg string) {
	res := &ResponseData{
		Status: status,
		Msg:    msg,
	}
	g.JSON(status, res)
}
