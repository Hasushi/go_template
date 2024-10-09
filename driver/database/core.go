package database

import (
	"fmt"
	"go_template/config"
	"log"
	"math"
	"os"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewMySQLDB(logger *zap.Logger, isLogging bool) (*gorm.DB, error) {
	dsn := config.DNS()

	gormConfig := &gorm.Config{}
	if !isLogging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} else {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Duration(config.SlowQueryThresholdMilliSecond()) * time.Millisecond,
				LogLevel:      gormLogger.Warn,
			},
		)
		gormConfig.Logger = newLogger
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.DBMaxIdleConns())
	sqlDB.SetMaxOpenConns(config.DBMaxOpenConns())

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		logger.Warn("failed to connect to database retrying", zap.Error(err), zap.Duration("sleepSeconds", duration))
		time.Sleep(duration)
	}

	return db, nil
}

func NewPostgreSQLDB(logger *zap.Logger, isLogging bool) (*gorm.DB, error) {
	dsn := config.DNS()

	gormConfig := &gorm.Config{}
	if !isLogging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} else {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Duration(config.SlowQueryThresholdMilliSecond()) * time.Millisecond,
				LogLevel:      gormLogger.Warn,
			},
		)
		gormConfig.Logger = newLogger
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open PostgreSQL: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.DBMaxIdleConns())
	sqlDB.SetMaxOpenConns(config.DBMaxOpenConns())

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		logger.Warn("failed to connect to database retrying", zap.Error(err), zap.Duration("sleepSeconds", duration))
		time.Sleep(duration)
	}

	return db, nil
}

func NewSQLiteDB(logger *zap.Logger, isLogging bool) (*gorm.DB, error) {
	dsn := config.DNS()

	gormConfig := &gorm.Config{}
	if !isLogging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} else {
		newLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Duration(config.SlowQueryThresholdMilliSecond()) * time.Millisecond,
				LogLevel:      gormLogger.Warn,
			},
		)
		gormConfig.Logger = newLogger
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(config.DBMaxIdleConns())
	sqlDB.SetMaxOpenConns(config.DBMaxOpenConns())

	// Check connection
	const retryMax = 10
	for i := 0; i < retryMax; i++ {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		if i == retryMax-1 {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		duration := time.Millisecond * time.Duration(math.Pow(1.5, float64(i))*1000)
		logger.Warn("failed to connect to database retrying", zap.Error(err), zap.Duration("sleepSeconds", duration))
		time.Sleep(duration)
	}

	return db, nil
}
