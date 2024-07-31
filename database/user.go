package database

import (
	"blog/util"

	"gorm.io/gorm"
)

type User struct { //默认情况下，驼峰对应的蛇形复数(users)即为数据库里的表名
	Id     int    `gorm:"column:id;primaryKey"`
	Name   string `gorm:"column:name"`     //name,可省略tag
	PassWd string `gorm:"column:password"` //pass_wd
}

func (User) TableName() string { //显示指定表名
	return "user"
}

var (
	_all_user_field = util.GetGormFields(User{})
)

// 根据用户名检索用户
func GetUserByName(name string) *User {
	db := GetBlogDBConnection()
	var user User
	// var users []User
	// select id,name form user where name="cpq" and password="123.com" order by id desc limit 20,4
	// db.Select([]string{"id", "name"}).Where("name='cpq' and password = '123.com'").Order("id desc").Offset(20).Limit(4).Find(&user)
	// 只查询一个条目：
	// db.Select("id,name").Where("name='cpq' and password = '123.com'").Order("id desc").Offset(20).Limit(4).First(&user)
	// db.Select("id,name").Where("name='" + name + "' and password = '123.com'").Order("id desc").Offset(20).Limit(4).Find(&users)
	// db.Select("id,name").Where("name=? and password = '123.com'", name).Order("id desc").Offset(20).Limit(4).Find(&users)
	if err := db.Select(_all_user_field).Where("name=?", name).First(&user).Error; err != nil {
		//name是unique的，所以用first即可
		if err != gorm.ErrRecordNotFound { //若是用户名不存在导致的错误，不需要打错误日志
			util.LogRus.Errorf("get password of user %s failed: %s", name, err) //系统性异常才打日志
		}
		return nil
	}
	return &user
}

func CreateUser(name, pass string) {
	db := GetBlogDBConnection()
	pass = util.Md5(pass)                          //密码经过md5
	user := User{Name: name, PassWd: pass}         //ORM
	if err := db.Create(&user).Error; err != nil { //必须传指针，因为要给user的主键赋值
		util.LogRus.Errorf("create user %s failed: %s", name, err)
	} else {
		util.LogRus.Errorf("create user id: %d", user.Id)
	}
}

// delete from user where name = "xxx"
func DeleteUser(name string) {
	db := GetBlogDBConnection()
	if err := db.Where("name=?", name).Delete(User{}).Error; err != nil {
		util.LogRus.Errorf("delete user %s failed: %s", name, err)
	}
}
