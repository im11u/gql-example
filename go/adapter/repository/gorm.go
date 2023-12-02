package repository

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// *sql.DBから*gorm.DBを取得する
func openGorm(sqlDB *sql.DB) *gorm.DB {
	mc := mysql.Config{
		Conn: sqlDB,
	}
	c := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(mysql.New(mc), &c)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
