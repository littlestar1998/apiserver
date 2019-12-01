package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	//"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
	}

	u := model.UserModel{
		//BaseModel: model.BaseModel{},
		Username: r.Username,
		Password: r.Password,
	}
	//fmt.Println()
	//log.Info(r.Username)
	//log.Info(r.Password)

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	SendResponse(c, nil, rsp)
}
