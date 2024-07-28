package database

import (
	"blog/util"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

var (
	blog_mysql      *gorm.DB
	blog_mysql_once sync.Once
	dblog           ormlog.Interface
	blog_redis      *redis.Client
	blog_redis_once sync.Once
)

func init() {
	dblog = ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		ormlog.Config{
			SlowThreshold: 100 * time.Millisecond, //慢SQL阈值
			LogLevel:      ormlog.Info,            //Log Level, Silent表示不输出日志
			Colorful:      true,                   //是否彩色打印日志
		},
	)
}

func createMysqlDB(dbname, host, user, pass string, port int) *gorm.DB {
	// dsn--data source name 是 blogtester@123.com@tcp(localhost:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass,
		host, port, dbname) //mb4兼容emoj表情
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dblog, PrepareStmt: true}) //启用PrepareStmt,SQL预编译,提高查询效率
	if err != nil {
		util.LogRus.Panicf("connect to mysql use dsn: %s failed %s", dsn, err)
	}
	//设置数据库连接池参数，提高并发性能
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	util.LogRus.Infof("connect to mysql db %s", dbname)
	return db
}

func GetBlogDBConnection() *gorm.DB { //连接池只需要创建一次，使用单例模式
	if blog_mysql == nil {
		dbname := "blog"
		viper := util.CreateConfig("mysql")
		host := viper.GetString(dbname + ".host")
		port := viper.GetInt(dbname + ".port")
		user := viper.GetString(dbname + ".user")
		pass := viper.GetString(dbname + ".pass")
		blog_mysql = createMysqlDB(dbname, host, user, pass, port)
	}

	return blog_mysql
}
