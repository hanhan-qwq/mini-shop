package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mini_shop/service"
	"net/http"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthService: service.NewAuthService(),
	}
}

type RegisterRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ctrl *AuthController) UserRegister(c *gin.Context) {
	//1.绑定参数
	req := RegisterRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
	}
	//2.校验两次密码是否一致
	if req.Password != req.ConfirmPassword {
		fmt.Println(req.ConfirmPassword)
		fmt.Println(req.Password)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "两次密码输入不一致",
		})
		return
	}
	//3.用户注册方法
	err := ctrl.AuthService.UserRegister(req.Username, req.Password, req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
	})
}

func (ctrl *AuthController) UserLogin(c *gin.Context) {
	//1.绑定参数
	req := LoginRequest{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "参数绑定失败",
			"error":   err.Error(),
		})
	}
	//2.校验用户信息
	accessToken, refreshToken, role, err := ctrl.AuthService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": "用户登录失败",
			"error":   err.Error(),
		})
	}
	//3.用户成功登录
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "用户成功登录",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"role":          role,
		},
	})
}
