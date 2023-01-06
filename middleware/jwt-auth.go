package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/linothomas14/hadir-in-api/helper"
	"github.com/linothomas14/hadir-in-api/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.BuildResponse("Use Bearer token please", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if authHeader == "" {
			response := helper.BuildResponse("No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := jwtService.ValidateToken(authHeader)

		if !token.Valid {
			log.Println(err)
			log.Println(token)
			response := helper.BuildResponse("Token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		log.Println("Claim[user_id]: ", claims["user_id"])
		log.Println("Claim[issuer] :", claims["iss"])
		c.Set("claims", claims)
		c.Set("user_id", claims["user_id"])
		c.Next()

	}

}
