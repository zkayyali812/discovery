package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToken
func GetToken(c *gin.Context) {
	token := c.PostForm("refresh_token")
	if token == "" {
		file, err := ioutil.ReadFile("data/auth_error.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusBadRequest, "application/json", file)
	} else {
		file, err := ioutil.ReadFile("data/auth_success.json")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": fmt.Sprintf("Error reading file: %s", err.Error()),
			})
			return
		}
		c.Data(http.StatusOK, "application/json", file)
	}
}