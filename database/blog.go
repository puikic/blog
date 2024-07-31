package database

import (
	"blog/util"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	Id         int       `gorm:"column:id;primaryKey"`
	UserId     int       `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Article    string    `gorm:"column:article"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (Blog) TableName() string { //显式指定表名
	return "blog"
}

var (
	_all_blog_field = util.GetGormFields(Blog{})
)

// 根据id获取博客内容
func GetBlogById(id int) *Blog {
	db := GetBlogDBConnection()
	var blog Blog
	if err := db.Select(_all_blog_field).Where("id=?", id).First(&blog).Error; err != nil { //必须传Blog的指针
		if err != gorm.ErrRecordNotFound {
			util.LogRus.Errorf("get content of blog %d failed: %s", id, err)
		}
		return nil
	}
	return &blog
}

// 根据作者id获取博客列表(仅包含博客id和标题)
func GetBlogByUserId(uid int) []*Blog {
	db := GetBlogDBConnection()
	var blogs []*Blog
	if err := db.Select("id,title").Where("user_id=?", uid).Find(&blogs).Error; err != nil { //必须传Blog的指针
		if err != gorm.ErrRecordNotFound {
			util.LogRus.Errorf("get blogs of user %d failed: %s", uid, err)
		}
		return nil
	}
	return blogs
}

// 根据博客id更新标题和正文
func UpdateBlogById(blog *Blog) error {
	if blog.Id <= 0 {
		return fmt.Errorf("couldn't update blog of id: %d", blog.Id) //errors.New("")
	}
	if len(blog.Title) == 0 || len(blog.Article) == 0 {
		return fmt.Errorf("could't set blog title or article to empty")
		// return errors.New("could't set blog title or article to empty")
	}
	db := GetBlogDBConnection()
	return db.Model(Blog{}).Where("id=?", blog.Id).Updates(map[string]any{"title": blog.Title, "article": blog.Article}).Error
}
