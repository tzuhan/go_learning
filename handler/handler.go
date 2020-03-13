package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
// @produce application/json
// @Success 200 string
// @Router /api/v1/role
func GetRoles(c *gin.Context) {
	message, err := json.Marshal(&Users)
	if err != nil {
		c.String(http.StatusOK, string(message))
		return
	}
	c.String(http.StatusInternalServerError, string(err.Error()))
	return
}

// @Summary Get role by id
// @Id getRole
// @Tags Role
// @version 1.0.0
// @produce application/json
// @Success 200 string
// @Router /api/v1/role/{id}
func GetRole(c *gin.Context) {
	for _, user := range Users {
		if id, _ := strconv.Atoi(c.Param("id")); id == user.Id {
			message, err := json.Marshal(&user)
			if err != nil {
				c.String(http.StatusOK, string(message))
				return
			} else {
				c.String(http.StatusInternalServerError, string(err.Error()))
				return
			}
		}
	}
	c.String(http.StatusBadRequest, "No such role id")
	return
}
