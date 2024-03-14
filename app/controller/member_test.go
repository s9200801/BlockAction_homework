package controller

import (
	"blockaction_homework/app/model"
	"blockaction_homework/app/service"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SuiteMemberController struct {
	suite.Suite
	controller  *MemberController
	mockService *mockMemberService
}

type mockMemberService struct {
	mock.Mock
}

func (m *mockMemberService) Register(param service.RegisterParam) (*model.Member, error) {
	args := m.Called(param)
	return args.Get(0).(*model.Member), args.Error(1)
}

func (m *mockMemberService) Login(param service.LoginParam) error {
	args := m.Called(param)
	return args.Error(0)
}

func (m *mockMemberService) SearchMember(param service.SearchMemberParam) (*model.Member, error) {
	args := m.Called(param)
	return args.Get(0).(*model.Member), args.Error(1)
}

func (m *mockMemberService) SearchAllMembers() ([]model.Member, error) {
	args := m.Called()
	return args.Get(0).([]model.Member), args.Error(1)
}

func (s *SuiteMemberController) SetupTest() {
	mockService := &mockMemberService{}
	controller := NewMemberController(mockService)
	s.controller = controller
	s.mockService = mockService
}

func (s *SuiteMemberController) Test_Register_success() {
	exceptMember := &model.Member{
		Id:       1,
		Username: "test",
		Password: "test",
		Mail:     "test",
	}
	exceptParm := service.RegisterParam{
		Username: "test",
		Password: "test",
		Mail:     "test",
	}

	s.mockService.On("Register", exceptParm).Return(exceptMember, nil)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	query, _ := json.Marshal(map[string]string{
		"username": "test",
		"password": "test",
		"mail":     "test",
	})
	engine.POST("/api/member", s.controller.Register)

	request, _ := http.NewRequest("POST", "/api/member", bytes.NewBuffer(query))
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	exceptResult, _ := json.Marshal(exceptMember)
	assert.Equal(s.T(), bytes.NewBuffer(exceptResult), responseRecorder.Body)
}

func (s *SuiteMemberController) Test_Register_fail() {
	exceptErr := errors.New("account is already registered")
	s.mockService.On("Register", mock.Anything).Return(&model.Member{}, exceptErr)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	query, _ := json.Marshal(map[string]string{
		"username": "test",
		"password": "test",
		"mail":     "test",
	})
	engine.POST("/api/member", s.controller.Register)

	request, _ := http.NewRequest("POST", "/api/member", bytes.NewBuffer(query))
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (s *SuiteMemberController) Test_Register_fail_missingField() {
	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	query, _ := json.Marshal(map[string]string{
		"username": "test",
		"password": "test",
	})
	engine.POST("/api/member", s.controller.Register)

	request, _ := http.NewRequest("POST", "/api/member", bytes.NewBuffer(query))
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (s *SuiteMemberController) Test_Login_success() {
	exceptParm := service.LoginParam{
		Username: "test",
		Password: "test",
	}

	s.mockService.On("Login", exceptParm).Return(nil)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	query, _ := json.Marshal(map[string]string{
		"username": "test",
		"password": "test",
	})
	engine.POST("/api/session", s.controller.Login)

	request, _ := http.NewRequest("POST", "/api/session", bytes.NewBuffer(query))
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
}

func (s *SuiteMemberController) Test_Login_fail() {
	exceptParm := service.LoginParam{
		Username: "test",
		Password: "test123",
	}
	exceptErr := errors.New("wrong password")

	s.mockService.On("Login", exceptParm).Return(exceptErr)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	query, _ := json.Marshal(map[string]string{
		"username": "test",
		"password": "test123",
	})
	engine.POST("/api/session", s.controller.Login)

	request, _ := http.NewRequest("POST", "/api/session", bytes.NewBuffer(query))
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (s *SuiteMemberController) Test_GetMemberData_success() {
	exceptParm := service.SearchMemberParam{
		Username: "test",
	}
	exceptMember := &model.Member{
		Id:       1,
		Username: "test",
		Mail:     "test",
	}

	s.mockService.On("SearchMember", exceptParm).Return(exceptMember, nil)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.GET("/api/member/:username", s.controller.GetMemberData)

	request, _ := http.NewRequest("GET", "/api/member/test", nil)
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	exceptResult, _ := json.Marshal(exceptMember)
	assert.Equal(s.T(), bytes.NewBuffer(exceptResult), responseRecorder.Body)
}

func (s *SuiteMemberController) Test_GetMemberData_fail() {
	exceptParm := service.SearchMemberParam{
		Username: "test",
	}
	exceptErr := errors.New("can not find username")
	s.mockService.On("SearchMember", exceptParm).Return(&model.Member{}, exceptErr)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.GET("/api/member/:username", s.controller.GetMemberData)

	request, _ := http.NewRequest("GET", "/api/member/test", nil)
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusNotFound, responseRecorder.Code)
}

func (s *SuiteMemberController) Test_GetAllMemberData_success() {
	exceptMembers := []model.Member{
		{
			Id:       1,
			Username: "test",
			Password: "test",
			Mail:     "test",
		},
		{
			Id:       2,
			Username: "test2",
			Password: "test2",
			Mail:     "test2",
		},
	}
	s.mockService.On("SearchAllMembers").Return(exceptMembers, nil)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.GET("/api/members", s.controller.GetAllMemberData)

	request, _ := http.NewRequest("GET", "/api/members", nil)
	request.Header.Add("Content-Type", "application/json")
	ctx.Request = request
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	exceptResult, _ := json.Marshal(exceptMembers)
	assert.Equal(s.T(), bytes.NewBuffer(exceptResult), responseRecorder.Body)
}

func TestSuiteMember(t *testing.T) {
	suite.Run(t, new(SuiteMemberController))
}
