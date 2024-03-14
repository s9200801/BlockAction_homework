package service

import (
	"blockaction_homework/app/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SuiteMemberService struct {
	suite.Suite
	db      *gorm.DB
	service MemberService
}

func InitDB() *gorm.DB {
	dns := "host=localhost user=postgres password=123 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	sqlDb, _ := gorm.Open(postgres.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "test.", // schema name
		},
	})
	sqlDb.Exec("CREATE SCHEMA IF NOT EXISTS \"test\";")
	return sqlDb
}

func (s *SuiteMemberService) SetupSuite() {
	db := InitDB()
	s.db = db
	service := NewMemberService(db)
	s.service = service
}

func (s *SuiteMemberService) SetupTest() {
	s.db.AutoMigrate(&model.Member{})
}

func (s *SuiteMemberService) TearDownTest() {
	s.db.Migrator().DropTable(&model.Member{})
}

func (s *SuiteMemberService) TearDownSuite() {
	s.db.Exec("DROP SCHEMA \"test\"")
}

func (s *SuiteMemberService) insertTestData() {
	s.db.Model(&model.Member{}).Create(&model.Member{
		Username: "test",
		Password: "test",
		Mail:     "test",
	})
}

func (s *SuiteMemberService) Test_Register_success() {
	registerParam := RegisterParam{
		Username: "test",
		Password: "test_password",
		Mail:     "test",
		Name:     "test",
		Phone:    "test",
	}
	result, err := s.service.Register(registerParam)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), registerParam.Username, result.Username)
	assert.Equal(s.T(), registerParam.Password, result.Password)
}

func (s *SuiteMemberService) Test_Register_fail_duplicateUserName() {
	s.insertTestData()

	registerParam := RegisterParam{
		Username: "test",
		Password: "test_password",
		Mail:     "test11",
		Name:     "test",
		Phone:    "test",
	}

	_, err := s.service.Register(registerParam)
	assert.EqualError(s.T(), err, "username is already registered")
}

func (s *SuiteMemberService) Test_Register_fail_duplicateMail() {
	s.insertTestData()

	registerParam := RegisterParam{
		Username: "test22",
		Password: "test_password",
		Mail:     "test",
		Name:     "test",
		Phone:    "test",
	}

	_, err := s.service.Register(registerParam)
	assert.EqualError(s.T(), err, "mail is already registered")
}

func (s *SuiteMemberService) Test_Login_success() {
	s.insertTestData()

	loginParam := LoginParam{
		Username: "test",
		Password: "test",
	}

	err := s.service.Login(loginParam)
	assert.Nil(s.T(), err)
}

func (s *SuiteMemberService) Test_Login_fail_notFoundUsername() {
	s.insertTestData()

	loginParam := LoginParam{
		Username: "test22",
		Password: "test",
	}

	err := s.service.Login(loginParam)
	assert.EqualError(s.T(), err, "not found username")
}

func (s *SuiteMemberService) Test_Login_fail_wrongPassword() {
	s.insertTestData()

	loginParam := LoginParam{
		Username: "test",
		Password: "test11",
	}

	err := s.service.Login(loginParam)
	assert.EqualError(s.T(), err, "wrong password")
}

func (s *SuiteMemberService) Test_SearchMember_success() {
	s.insertTestData()

	searchMemberParam := SearchMemberParam{
		Username: "test",
	}

	result, err := s.service.SearchMember(searchMemberParam)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), searchMemberParam.Username, result.Username)
}

func (s *SuiteMemberService) Test_SearchMember_fail_notFound() {
	s.insertTestData()

	searchMemberParam := SearchMemberParam{
		Username: "test11",
	}

	_, err := s.service.SearchMember(searchMemberParam)
	assert.EqualError(s.T(), err, "can not find member")
}

func (s *SuiteMemberService) Test_SearchAllMember_success() {
	s.insertTestData()

	members, err := s.service.SearchAllMembers()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(members))
	assert.Equal(s.T(), "test", members[0].Username)
}

func (s *SuiteMemberService) Test_SearchAllMember_success_emptyList() {
	members, err := s.service.SearchAllMembers()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 0, len(members))
}

func TestSuiteMember(t *testing.T) {
	suite.Run(t, new(SuiteMemberService))
}
