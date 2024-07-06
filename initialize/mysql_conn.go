package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"xx/global"
)

type dbConfig struct {
	dbType      string //数据库类型 mysql
	dbName      string //数据库名称
	username    string //数据库账号/用户名
	password    string //数据库密码
	host        string //数据库ip 域名
	port        int
	tablePrefix string //表头
}

func init() {
	fmt.Println("initMysql")
	var err error
	cfg := &dbConfig{
		dbType:   global.ServerConfig.MysqlInfo.DbType,
		dbName:   global.ServerConfig.MysqlInfo.DbName,
		username: global.ServerConfig.MysqlInfo.Username,
		password: global.ServerConfig.MysqlInfo.Password,
		host:     global.ServerConfig.MysqlInfo.Host,
		port:     global.ServerConfig.MysqlInfo.Port,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.dbName)

	global.MysqlConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql连接错误:", err)
	} else {
		fmt.Println("mysql连接成功")
	}
}
