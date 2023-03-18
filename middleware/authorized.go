package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"fake.com/medium/models"
	"fake.com/medium/policies"
	"fake.com/medium/response"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware politica de acceso de  recurso de solicitud solo si se encuentra asignado

func AuthorizedMiddleware(authorized policies.Authorized) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario models.Usuario
		usernameInterface, _ := c.Get(authorizationUsername)
		username := fmt.Sprintf("%v", usernameInterface)
		usuario = authorized.Me(username)

		if usuario.Rol == "ADMIN" {
			c.Next()
			return
		}

		c.Next()
	}
}

func acceso(c *gin.Context, accede bool) {
	if accede == false {
		err := errors.New("No cuenta con autorización para este servicio ")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}
	c.Next()
	return
}
