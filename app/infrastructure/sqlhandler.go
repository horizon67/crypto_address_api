package infrastructure

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	"github.com/horizon67/crypto_address_api/app/domain"
	"github.com/horizon67/crypto_address_api/app/interfaces/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := gorm.Open("mysql", "root:root@tcp(db)/crypto_address_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error)
	}
	conn.LogMode(true)

	m := gormigrate.New(conn, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create crypto_addresses table
		{
			ID: "202006090000",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.CryptoAddress{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("crypto_address").Error
			},
		},
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
