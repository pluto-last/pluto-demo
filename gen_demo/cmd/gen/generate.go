package main

// gorm gen configure

import (
	"gen_demo/model/table"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"

	"gorm.io/gen"
)

var MyDBDSN = "host=127.0.0.1 user=pluto dbname=pluto sslmode=disable password=880206congwu"

func main() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		// 默认会在 OutPath 目录生成CRUD代码，并且同目录下生成 model 包
		// 所以OutPath最终package不能设置为model，在有数据库表同步的情况下会产生冲突
		// 若一定要使用可以通过ModelPkgPath单独指定model package的名称
		OutPath: "../../dal/query",
		/* ModelPkgPath: "dal/model"*/

		// gen.WithoutContext：禁用WithContext模式
		// gen.WithDefaultQuery：生成一个全局Query对象Q
		// gen.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(Instance(MyDBDSN))

	// 从连接的数据库为所有表生成Model结构体和CRUD代码
	// 也可以手动指定需要生成代码的数据表
	g.ApplyBasic(table.XcData{})

	// 执行并生成代码
	g.Execute()
}

var once = new(sync.Once)

func Instance(dsn string) *gorm.DB {

	// 默认的日志对象
	var ormLogger logger.Interface

	ormLogger = logger.Default

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		return nil
	}

	//// 设置批量插入的规模
	//db = db.Session(&gorm.Session{CreateBatchSize: 1000})

	sqlDB, _ := db.DB()

	if err := sqlDB.Ping(); err != nil {
		return nil
	}

	// 设置连接池参数
	once.Do(func() {
		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(200)
		sqlDB.SetConnMaxIdleTime(30 * time.Second)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db
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
