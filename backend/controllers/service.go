package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rathink4/backend/initializers"
	"github.com/rathink4/backend/utils"
)

func OAuthLogin(c *gin.Context) {
	url := initializers.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")

	c.Redirect(http.StatusSeeOther, url)
}

func OAuth(c *gin.Context) {
	state := c.Query("state")

	if state != "randomstate" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "States do not match",
		})
		return
	}

	// Get the code
	code := c.Query("code")

	googleOAuthConfig := initializers.GoogleConfig()

	// Handle the exchange code to initiate a transport.
	tok, err := googleOAuthConfig.Exchange(context.Background(), code)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// Get user detail
	userInfo, err := utils.GetUserInfo(tok.AccessToken)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, userInfo)
}
