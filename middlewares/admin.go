package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/utils"
)

func AdminAuthMiddleware(c *gin.Context) {
	Token, err := c.Cookie("Authorisation")
	if err != nil {
		fmt.Println("he")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	role, err := utils.GetRoleFromToken(Token)
	if err != nil {
		fmt.Println("ji")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if role == "admin" {
		c.Next()
	}else{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	
}

// func AdminAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenHeader := c.GetHeader("authorization")
// 		fmt.Println(tokenHeader, "this is the token header")
// 		if tokenHeader == "" {
// 			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return
// 		}

// 		splitted := strings.Split(tokenHeader, " ")
// 		if len(splitted) != 2 {
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return

// 		}
// 		tokenpart := splitted[1]
// 		tokenClaims, err := helper.ValidateToken(tokenpart)
// 		if err != nil {
// 			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token  ", nil, err.Error())
// 			c.JSON(http.StatusUnauthorized, response)
// 			c.Abort()
// 			return

// 		}
// 		c.Set("tokenClaims", tokenClaims)

// 		c.Next()

// 	}

// }
