package repository

import (
	"log"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {

	//dsn := os.Getenv("DB_CONNECTION")
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	xraycon, err := xray.SQLContext(os.Getenv("XRAY_DB_TYPE"), os.Getenv("XRAY_DB_CONNECTION"))
	if err != nil {
		log.Println(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: xraycon,
	}), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	return db
}
