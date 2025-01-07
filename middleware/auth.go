package middleware

import (
	"fmt"
	"net/http"

	"github.com/DevisArya/reservasi_lapangan/utils"
	"github.com/labstack/echo/v4"
)

func Auth(role []string, checkUserId bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			claims, err := utils.VerifyToken(c)
			if err != nil {
				return err
			}

			//jika perlu mengecek role untuk akses
			if len(role) != 0 {
				userRole, ok := (*claims)["role"].(string)

				if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{
						"message": "Role is missing in token",
					})
				}

				found := false
				for _, val := range role {
					if userRole == val {
						found = true
						break
					}
				}

				if !found {
					return c.JSON(http.StatusForbidden, map[string]string{
						"message": "Forbidden: You don't have access to this resource",
					})
				}
			}

			//jika perlu mengecek id user untuk akses
			if checkUserId {
				userIdInterface, ok := (*claims)["user_id"]
				if !ok {
					return c.JSON(http.StatusUnauthorized, map[string]string{
						"message": "User ID is missing in token",
					})
				}

				userId := fmt.Sprintf("%v", userIdInterface)
				userIdParam := c.Param("id")

				if userId != userIdParam {
					return c.JSON(http.StatusUnauthorized, map[string]string{
						"message": "User ID does not match",
					})
				}
			}

			c.Set("claims", claims)

			return next(c)
		}
	}
}
