package database

import (
	"github.com/jinzhu/gorm"
)

type IDbProvider interface {
	Instance() *gorm.DB
	Ping() error
}
