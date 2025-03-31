package mysql

import "gorm.io/gorm"

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *mysqlRepository {
	return &mysqlRepository{db: db}
}
