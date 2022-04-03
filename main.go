package main

import (
	handler "functionX-Query-APIs/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/query/community-pool/outstanding", handler.CommunityPollOutstanding)
	router.GET("/query/community-pool/deductions", handler.CommunityPollDeductions)

	router.Run(":5000")

}
