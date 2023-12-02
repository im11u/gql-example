package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	c := newConfig()

	mc := mysql.Config{
		Net:       "tcp",
		Addr:      c.host + ":" + c.port,
		DBName:    c.database,
		User:      c.user,
		Passwd:    c.password,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
		ParseTime: true,
	}

	return sql.Open("mysql", mc.FormatDSN())
}
