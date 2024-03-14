package controller

import (
	"blockaction_homework/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
	MemberService service.MemberServiceInterface
}

func NewMemberController(memberService service.MemberServiceInterface) *MemberController {
	return &MemberController{
		MemberService: memberService,
	}
}

// 會員註冊
func (controller MemberController) Register(c *gin.Context) {
	var registerParam service.RegisterParam
	if err := c.ShouldBind(&registerParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
		return
	}
	if registerMember, err := controller.MemberService.Register(registerParam); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"msg": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, registerMember)
	}
}

// 會員登入
func (controller MemberController) Login(c *gin.Context) {
	var loginParam service.LoginParam
	if err := c.ShouldBind(&loginParam); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if err := controller.MemberService.Login(loginParam); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

// 取得會員資料
func (controller MemberController) GetMemberData(c *gin.Context) {
	username := c.Param("username")
	searchMemberParam := service.SearchMemberParam{
		Username: username,
	}
	member, err := controller.MemberService.SearchMember(searchMemberParam)
	if err != nil {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, member)
	}
}

// 取得所有會員資料
func (controller MemberController) GetAllMemberData(c *gin.Context) {
	members, err := controller.MemberService.SearchAllMembers()
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, members)
	}
}
