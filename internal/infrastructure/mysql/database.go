package mysql

import (
	"backend/internal/entities"
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

func (mysql *dbClient) Add(product entities.Product) error {
	result := mysql.client.Create(&product)
	return result.Error
}

func (mysql *dbClient) Get(productID int) (entities.Product, error) {
	var product entities.Product
	result := mysql.client.First(&product, productID)
	return product, result.Error
}

func (mysql *dbClient) GetAll() ([]entities.Product, error) {
	var products []entities.Product
	result := mysql.client.Find(&products)
	return products, result.Error
}
