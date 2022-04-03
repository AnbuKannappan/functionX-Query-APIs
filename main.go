package main

import (
	"functionX-Query-APIs/environments"
	handler "functionX-Query-APIs/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	environment := new(environments.FXCORE_ENV)
	err := environment.Load()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/query/community-pool/outstanding", handler.CommunityPollOutstanding)
	router.GET("/query/community-pool/deductions", handler.CommunityPollDeductions)

	router.Run(environment.App_Port)

}
