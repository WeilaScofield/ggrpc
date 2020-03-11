package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func NewConn(host string, port int, user string, passwd string, dbname string) *gorm.DB {
	conf := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, passwd)
	db, err := gorm.Open("postgres", conf)
	if err != nil {
		log.Fatal("Postgres init failed:", err)
	}
	db.DB().SetMaxIdleConns(5)  //最大闲置连接
	db.DB().SetMaxOpenConns(10) //最大连接
	return db
}
