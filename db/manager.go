package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Addr     string `json:"server" env:"DB_ADDR" envDefault:"127.0.0.1"`
	Port     string `json:"port" env:"DB_PORT" envDefault:"5432"`
	Database string `json:"database" env:"DB_USER" envDefault:"user"` //TODO: to do for production
	Username string `json:"username" env:"DB_USER" envDefault:"user"`
	Password string `json:"password" env:"DB_PASSWORD" envDefault:"password"`

	MaxIdleConns    int `json:"maxidleconns" env:"MAX_IDLE_CONNS" envDefault:"5"`
	MaxOpenConns    int `json:"maxopenconns" env:"MAX_OPEN_CONNS" envDefault:"5"`
	ConnMaxLifetime int `json:"connmaxlifetime" env:"CONN_MAX_LIFETIME" envDefault:"90"`
}

func connectPostgres(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
		cfg.Addr, cfg.Port, cfg.Username, cfg.Password, cfg.Database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// }) //TODO
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GormOpen(cfg *Config) (*gorm.DB, error) {

	db, err := connectPostgres(cfg)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return db, nil
}
