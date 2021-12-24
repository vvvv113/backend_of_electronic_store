package mysql

import (
	"backend/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbClient struct {
	client *gorm.DB
}

type Config struct {
	IP       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func New(cnfg *Config) (*dbClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cnfg.User, cnfg.Password, cnfg.IP, cnfg.Port, cnfg.Database)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Critical.Fatalf("Failed to conect to Mongo %v", err)
	}
	return &dbClient{
		client: conn,
	}, nil
}

func (mysql *dbClient) Create(obj interface{}) error {
	result := mysql.client.Create(obj)
	return result.Error
}

func (mysql *dbClient) FindAll(obj interface{}) error {
	result := mysql.client.Find(obj)
	return result.Error
}

func (mysql *dbClient) FindByParameters(searchObj interface{}, obj interface{}, isAll bool) error {
	var result *gorm.DB
	if isAll {
		result = mysql.client.Where(searchObj).Find(obj)
	} else {
		result = mysql.client.Where(searchObj).First(obj)
	}
	return result.Error
}

func (mysql *dbClient) FindByID(ID int, obj interface{}) error {
	result := mysql.client.Where("id = ?", ID).First(obj)
	return result.Error
}

func (mysql *dbClient) Update(obj interface{}, key string, value string) error {
	result := mysql.client.Model(obj).Update(key, value)
	return result.Error
}
