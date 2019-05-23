package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

const pass  = "mr"

type Admin struct {
	Id int
	Username string `orm:"size(64) index" `
	Password string `orm:"size(64)"`
	Salt string `orm:"size(64)"`
	Created time.Time
}

// 通过用户名获取账号信息
func (this *Admin) GetByName(username string) *Admin {

	query := orm.NewOrm().QueryTable(this)
	_ = query.Filter("username", username).One(this)

	return this
}

// 检查密码是否正确
func (this *Admin) CheckPwd (pwd string) bool {
	return strings.EqualFold(this.Encode(pwd),this.Password)
}

// 加密密码
func (this *Admin) Encode(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd + pass + this.Salt))
	return hex.EncodeToString(h.Sum(nil))
}
