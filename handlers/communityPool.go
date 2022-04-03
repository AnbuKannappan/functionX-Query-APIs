package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

func CommunityPollOutstanding(c *gin.Context) {

	cmd := exec.Command("fxcored", "query", "distribution", "community-pool", "--node", "https://fx-json.functionx.io:26657")

	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, string(stdout))

}

func CommunityPollDeductions(c *gin.Context) {

	cmd := exec.Command("fxcored", "q", "gov", "proposals", "--node", "https://fx-json.functionx.io:26657")

	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, string(stdout))

}
