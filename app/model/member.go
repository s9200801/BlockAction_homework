package model

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	Id         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username   string    `json:"username" form:"username" gorm:"column:username;uniqueIndex;not null"`
	Password   string    `json:"-" form:"password" gorm:"column:password;not null"`
	Mail       string    `json:"mail" form:"mail" gorm:"mail;not null"`
	Name       string    `json:"name" form:"name" gorm:"name"`
	Phone      string    `json:"phone" form:"phone" gorm:"phone"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time;autoCreateTime"`
}

// 初始化建立表單
func (m Member) InitTable(db *gorm.DB) {
	db.AutoMigrate(&m)
}

// 建立會員資料
func CreateMember(db *gorm.DB, registerMember *Member) error {
	// SQL 語法
	// INSERT INTO members (username, password, name, phone, mail)
	// VALUES (username, password, name, phone, mail);
	err := db.Model(&Member{}).Create(registerMember).Error
	return err
}

// 撈取單一會員資料By UserName
func GetMemberByUserName(db *gorm.DB, userName string) (Member, error) {
	var member Member
	// SQL 語法
	// SELECT * FROM members
	// WHERE userName = userName
	// LIMIT 1;
	err := db.Model(&Member{}).Where("username = ?", userName).Take(&member).Error
	return member, err
}

//撈取單一會員資料By Mail
func GetMemberByMail(db *gorm.DB, mail string) (Member, error) {
	var member Member
	// SQL 語法
	// SELECT * FROM members
	// WHERE mail = mail
	// LIMIT 1
	err := db.Model(&Member{}).Where("mail = ?", mail).Take(&member).Error
	return member, err
}

// 撈取所有會員資料
func GetAllMember(db *gorm.DB) ([]Member, error) {
	var members []Member
	// SQL 語法
	// SELECT * FROM members
	// LIMIT 1
	err := db.Model(&Member{}).Find(&members).Error
	return members, err
}
