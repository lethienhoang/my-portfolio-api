package https

import "github.com/gin-gonic/gin"

// GetClientIP gets a requests IP address by reading off the forwarded-for
func GetClientIP(c *gin.Context) string {
	forwarded := c.Request.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}

	return c.Request.RemoteAddr
}

// GetUserID gets the current_user ID as a string
func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if exists {
		return userID.(string)
	}

	return ""
}
