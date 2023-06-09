package db

import (
	"gV2/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var once = new(sync.Once)

func Instance(cfg *config.Specification) (*gorm.DB, error) {

	// 默认的日志对象
	var ormLogger logger.Interface
	if cfg.DB.Debug {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(postgres.Open(cfg.DB.DSN), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		return nil, err
	}

	//// 设置批量插入的规模
	//db = db.Session(&gorm.Session{CreateBatchSize: 1000})

	sqlDB, _ := db.DB()

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// 设置连接池参数
	once.Do(func() {
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(200)
		sqlDB.SetConnMaxIdleTime(30 * time.Second)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db, nil
}

type Tx struct {
	*gorm.DB
	commit bool
}

func Begin(db *gorm.DB) *Tx {
	return &Tx{
		DB: db.Begin(),
	}
}

func (tx *Tx) RollbackIfFailed() {
	if tx.commit {
		return
	}
	if err := tx.Rollback().Error; err != nil {
		log.Println("rollback failed", err)
	}
}

func (tx *Tx) Commit() {
	if err := tx.DB.Commit().Error; err != nil {
		log.Println("commit failed", err)
		return
	}
	tx.commit = true
}

func WithUpdatedAt(kv ...interface{}) map[string]interface{} {
	var k string
	m := make(map[string]interface{})
	for i, v := range kv {
		if i%2 != 0 {
			m[k] = v
			continue
		}
		k = v.(string)
	}
	m["updated_at"] = time.Now()
	return m
}
