package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var Users []User = []User{
	User{Id: 1, Name: "Admin"},
	User{Id: 2, Name: "Moderator"},
	User{Id: 3, Name: "Viewer"},
}

// @Summary Get all roles
// @Id getRoles
// @Tags Role
// @version 1.0.0
// @produce json
// @Success 200 {string} User
// @Failure 400 {} string
// @Failure 404 {} string
// @Failure 500 {} string
// @Router /role [get]
func GetRoles(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
	return
}
