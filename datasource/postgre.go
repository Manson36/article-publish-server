package datasource

import (
	"fmt"
	"github.com/article-publish-server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type pqDB struct {
	*gorm.DB
}

func (p *pqDB) initDB() {

	arg := config.Postgres.GetURI()
	fmt.Println(arg)
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		msg := fmt.Sprintf("init postgres db error, msg:%s", err.Error())
		panic(msg)
	}

	if err = db.DB().Ping(); err != nil {
		msg := fmt.Sprintf("ping postgre db error, msg:%s", err.Error())
		panic(msg)
	}

	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.DB().SetMaxOpenConns(30)

	p.DB = db
}

var PqDB = &pqDB{}
