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
// @produce json
// @Success 200 {array} User
// @Failure 400 {} string
// @Failure 404 {} string
// @Failure 500 {} string
// @Router /role [get]
func GetRoles(c *gin.Context) {
	message, err := json.Marshal(&Users)
	if err != nil {
		c.String(http.StatusInternalServerError, "Json Parse Error.")
		return
	}
	c.String(http.StatusOK, string(message))
	return
}

// @Summary Get one role based on ID
// @Id getRole
// @Tags Role
// @version 1.0.0
// @produce json
// @Success 200 {object} User
// @Failure 400 {} string
// @Failure 404 {} string
// @Failure 500 {} string
// @Router /role/:id [get]
func GetRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Id invalid.")
		return
	}
	ind := -1
	for i := 0; i < len(Users); i++ {
		if Users[i].Id == id {
			ind = i
			break
		}
	}
	if ind == -1 {
		c.String(http.StatusBadRequest, "No user with this id.")
	} else {
		message, _ := json.Marshal(&Users[ind])
		c.String(http.StatusOK, string(message))
	}
	return

}
