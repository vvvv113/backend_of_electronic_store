package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}

type Config struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func New(cnfg Config) (*dbClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cnfg.User, cnfg.Password, cnfg.IP, cnfg.Port, cnfg.Database)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

	}

	return &dbClient{
		client: conn,
	}, nil
}
