package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// MatchUserTypeToUid 确保只有管理员才能访问用户数据
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil
	// type 是用户时，并且用户id不相等（表示该用户不是访问自己的数据）时，报错。
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}
