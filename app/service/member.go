package service

import (
	"blockaction_homework/app/model"

	"errors"

	"gorm.io/gorm"
)

type MemberServiceInterface interface {
	Register(RegisterParam) (*model.Member, error)
	Login(LoginParam) error
	SearchMember(SearchMemberParam) (*model.Member, error)
	SearchAllMembers() ([]model.Member, error)
}

type MemberService struct {
	baseService
}

type RegisterParam struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Mail     string `form:"mail"  binding:"required"`
	Name     string `form:"name"`
	Phone    string `form:"phone"`
}

type LoginParam struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SearchMemberParam struct {
	Username string
}

func NewMemberService(db *gorm.DB) MemberService {
	return MemberService{
		baseService: newBaseService(db),
	}
}

// 會員註冊
func (memberService MemberService) Register(registerParam RegisterParam) (*model.Member, error) {
	if _, err := model.GetMemberByUserName(memberService.DB, registerParam.Username); err == nil {
		return nil, errors.New("username is already registered")
	}
	if _, err := model.GetMemberByMail(memberService.DB, registerParam.Mail); err == nil {
		return nil, errors.New("mail is already registered")
	}
	registerMember := &model.Member{
		Username: registerParam.Username,
		Password: registerParam.Password,
		Name:     registerParam.Name,
		Phone:    registerParam.Phone,
		Mail:     registerParam.Mail,
	}
	if err := model.CreateMember(memberService.DB, registerMember); err != nil {
		panic(err)
	} else {
		return registerMember, nil
	}
}

// 會員登入
func (memberService MemberService) Login(loginParam LoginParam) error {
	member, err := model.GetMemberByUserName(memberService.DB, loginParam.Username)
	if err != nil {
		return errors.New("not found username")
	}
	if loginParam.Password != member.Password {
		return errors.New("wrong password")
	}
	return nil
}

// 搜尋會員
func (memberService MemberService) SearchMember(searchMemberParam SearchMemberParam) (*model.Member, error) {
	member, err := model.GetMemberByUserName(memberService.DB, searchMemberParam.Username)
	if err != nil {
		return nil, errors.New("can not find member")
	} else {
		return &member, nil
	}
}

// 搜尋所有會員
func (memberService MemberService) SearchAllMembers() ([]model.Member, error) {
	members, err := model.GetAllMember(memberService.DB)
	if err != nil {
		panic(err)
	} else {
		return members, nil
	}
}
