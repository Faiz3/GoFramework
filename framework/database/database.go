package database

import (
	"fmt"
	"sync"
	"time"

	"go-framework/framework/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	config config.DatabaseConfig
	db     *gorm.DB
	once   sync.Once
	err    error
}

var instance *Database

func New(cfg *config.Config) *Database {
	d := &Database{
		config: cfg.GetDatabaseConfig(),
	}
	instance = d
	return d
}

func (d *Database) connect() {
	d.once.Do(func() {
		var dialector gorm.Dialector
		switch d.config.Driver {
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				d.config.Username, d.config.Password, d.config.Host, d.config.Port, d.config.Database)
			dialector = mysql.Open(dsn)
		case "postgres":
			dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				d.config.Host, d.config.Port, d.config.Username, d.config.Password, d.config.Database)
			dialector = postgres.Open(dsn)
		case "sqlite":
			dialector = sqlite.Open(d.config.Database + ".db")
		default:
			d.err = fmt.Errorf("unsupported database driver: %s", d.config.Driver)
			return
		}

		db, err := gorm.Open(dialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			d.err = fmt.Errorf("failed to connect to database: %w", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			d.err = fmt.Errorf("failed to get underlying DB: %w", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		d.db = db
	})
}

func GetInstance() *Database {
	return instance
}

func (d *Database) GetDB() (*gorm.DB, error) {
	d.connect()
	if d.err != nil {
		return nil, d.err
	}
	return d.db, nil
}

func (d *Database) AutoMigrate(models ...interface{}) error {
	db, err := d.GetDB()
	if err != nil {
		return err
	}
	return db.AutoMigrate(models...)
}

func (d *Database) Close() error {
	if d.db == nil {
		return nil
	}
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func DB() (*gorm.DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("database not initialized")
	}
	return instance.GetDB()
}
