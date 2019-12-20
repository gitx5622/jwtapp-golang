package Models

import (
	"jwtapp/Config"
	"jwtapp/Services"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Authority struct {
	gorm.Model
	AuthorityId   string   `json:"authorityId" gorm:"not null;unique"`
	AuthorityName string `json:"authorityName"`
}

func (Authority) TableName() string {
	return "authorities"
}

func (a *Authority) CreateAuthority() (err error, authority *Authority) {
	err = Config.DB.Create(a).Error
	return err, a
}


func (a *Authority) DeleteAuthority() (err error) {
	err = Config.DB.Where("authority_id = ?", a.AuthorityId).Find(&User{}).Error
	if err != nil {
		err = Config.DB.Where("authority_id = ?", a.AuthorityId).Delete(a).Error
	} else {
		err = errors.New("There is an error")
	}
	return err
}


func (a *Authority) GetInfoList(info Services.PageInfo) (err error, list interface{}, total int) {

	err, total = Services.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var authority []Authority
		err = Config.DB.Find(&authority).Error
		return err, authority, total
	}
}
