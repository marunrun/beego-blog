package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id         int
	Title      string
	Created    time.Time `orm:"auto_now_add;index"`
	Views      int64     `orm:"index"`
	TopicCount int64
	Topics      []*Topic `orm:"reverse(many)"`
}

type Topic struct {
	Id              uint
	Category        *Category `orm:"column(category_id);rel(fk)"`
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int
}

func RegisterDB() {
	orm.RegisterModel(new(Category), new(Topic), new(Admin))
	orm.RegisterDataBase("default", beego.AppConfig.String("db_driver"), beego.AppConfig.String("db_connection"), 10)
}
