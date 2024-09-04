package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func link_test_db(path string) {

	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err))
	}

}

type Game struct {
	gorm.Model

	// 倒数多少秒
	Countdown int
	// 持续时间
	Duration int

	Status int
}

type Top struct {
	gorm.Model

	// 用户 id
	Userid string
	// 次数
	Times int

	GameId uint

	// 有效成绩
	Enable bool
}

func init() {
	path := "./tmp/db.db"
	link_test_db(path)

	DB.AutoMigrate(&Game{})
	DB.AutoMigrate(&Top{})
}
