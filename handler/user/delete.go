package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create creates a new user account.
func Delete(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.Atoi(id)
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
