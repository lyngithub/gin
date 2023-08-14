package models

import (
	"fmt"
	"xx/global"
)

type User struct {
	Id            int `json:"id" gorm:"column:id"`
	Online        int `json:"online" gorm:"column:online"`
	LastLoginTime int `json:"last_login_time" gorm:"column:last_login_time"`
}

func (t *User) TableNam() string {
	return "cmf_user"
}

func (t *User) Update() error {

	tx := global.MysqlConn.Table(t.TableNam()).Where("id=?", t.Id).Updates(map[string]interface{}{"online": t.Online, "last_login_time": t.LastLoginTime})
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return tx.Error
	}
	return nil
}
