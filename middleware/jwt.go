package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iskhakmuhamad/teaservice/shared"
	"github.com/iskhakmuhamad/teaservice/usecases"
)

func AuthorizeJWT(u usecases.Token) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response := shared.BuildErrorResponse("Failed", "No token found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := u.ValidateToken(authHeader)
		if token.Valid {
			_ = token.Claims.(jwt.MapClaims)
		} else {
			response := shared.BuildErrorResponse("Token is not valid", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
