package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

func InitGormDB() (err error) {
	dB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      "wechat:Kewa6BrXBFB63Hc7@tcp(42.192.155.29:3306)/wechat?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:        171,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Printf("连接失败：%v\n", err)
	}
	db = dB
	return err
}
