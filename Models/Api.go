package Models

import (
	"jwtapp/Config"
	"jwtapp/Services"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

)

type Api struct {
	gorm.Model
	Path        string `json:"path"`
	Description string `json:"description"`
	Group       string `json:"group"`
}

func (a *Api) CreateApi() (err error) {
	findOne := Config.DB.Where("path = ?", a.Path).Find(&Menu{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = Config.DB.Create(a).Error
	}
	return err
}

func (a *Api) DeleteApi() (err error) {
	err = Config.DB.Delete(a).Error
	err = Config.DB.Where("api_id = ?", a.ID).Unscoped().Delete(&ApiAuthority{}).Error
	return err
}

func (a *Api) UpdataApi() (err error) {
	err = Config.DB.Save(a).Error
	return err
}

func (a *Api) GetApiById(id float64) (err error, api Api) {
	err = Config.DB.Where("id = ?", id).First(&api).Error
	return
}


func (a *Api) GetAllApis() (err error, apis []Api) {
	err = Config.DB.Find(&apis).Error
	return
}


func (a *Api) GetInfoList(info Services.PageInfo) (err error, list interface{}, total int) {

	err, total = Services.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []Api
		err = Config.DB.Order("group", true).Where("path LIKE ?", "%"+a.Path+"%").Find(&apiList).Error
		return err, apiList, total
	}
}
