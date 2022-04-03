package handlers

import (
	utils "functionX-Query-APIs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func CommunityPollOutstanding(c *gin.Context) {

	cmd := exec.Command("fxcored", "query", "distribution", "community-pool", "--node", "https://fx-json.functionx.io:26657")

	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res := utils.YamlParser(string(stdout))

	c.JSON(http.StatusOK, res)

}

func CommunityPollDeductions(c *gin.Context) {

	cmd := exec.Command("fxcored", "q", "gov", "proposals", "--node", "https://fx-json.functionx.io:26657")

	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res := utils.YamlParser(string(stdout))

	c.JSON(http.StatusOK, res)

}
