package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

var mysql_pool sync.Pool

//mysql创建连接，和放入golang内置的链接池中
func CreateMysqlConnection(user, passwd, dbname string) {

	mysql_pool.New = func() interface{} {

		db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, passwd, dbname))
		if err != nil {
			log.Fatalln("链接mysql出错：", err)
		}

		return db
	}

}

//获取mysql实力,线程安全的,会自动释放
func GetMysqlDB() *gorm.DB {
	mysql := mysql_pool.Get()
	return mysql.(*gorm.DB)
}
