package controller

import (
	"github.com/gin-gonic/gin"
	"mini_shop/service"
	"net/http"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (ctrl *UserController) GetUserProfile(c *gin.Context) {
	//1.从jwt中间件获取userID
	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  "用户ID格式错误",
		})
		return
	}
	//2.获取用户信息
	s := service.NewUserService()
	profile, err := s.GetProfileByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  -1,
			"msg":   "获取用户信息失败",
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询用户数据成功",
		"data":    profile,
	})
}

func (ctrl *UserController) UpdateUserProfile(c *gin.Context) {
	//1.绑定参数
	req := UpdateProfileRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
	}
	//2.从jwt中间件获取userID
	userIDAny, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    -1,
			"message": "未登录",
		})
		return
	}
	userID, ok := userIDAny.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户ID格式错误",
		})
		return
	}
	//3.更新用户数据
	s := service.NewUserService()
	err := s.UpdateUserProfile(userID, req.Username, req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户信息更新失败",
			"error":   err.Error(),
		})
	}
}
