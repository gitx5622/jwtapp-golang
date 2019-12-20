package Models

import (
	"fmt"
	"jwtapp/Config"
	"jwtapp/Services"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

)

type BaseMenu struct {
	gorm.Model
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Meta      `json:"meta"`
	NickName  string     `json:"nickName"`
	Children  []BaseMenu `json:"children"`
}

func (b *BaseMenu) AddBaseMenu() (err error) {
	findOne := Config.DB.Where("name = ?", b.Name).Find(&BaseMenu{}).Error
	if findOne != nil {
		b.NickName = b.Title
		err = Config.DB.Create(b).Error
	} else {
		err = errors.New("存在重复name，请修改name")
	}
	return err
}

func (b *BaseMenu) DeleteBaseMenu(id float64) (err error) {
	err = Config.DB.Where("parent_id = ?", id).First(&BaseMenu{}).Error
	if err != nil {
		err = Config.DB.Where("id = ?", id).Delete(&b).Error
		err = Config.DB.Where("menu_id = ?", id).Unscoped().Delete(&Menu{}).Error
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

func (b *BaseMenu) UpdataBaseMenu() (err error) {
	upDataMap := make(map[string]interface{})
	upDataMap["parent_id"] = b.ParentId
	upDataMap["path"] = b.Path
	upDataMap["name"] = b.Name
	upDataMap["hidden"] = b.Hidden
	upDataMap["component"] = b.Component
	upDataMap["title"] = b.Title
	upDataMap["icon"] = b.Icon
	err = Config.DB.Where("id = ?", b.ID).Find(&BaseMenu{}).Updates(upDataMap).Error
	err1 := Config.DB.Where("menu_id = ?", b.ID).Find(&[]Menu{}).Updates(upDataMap).Error
	fmt.Printf("菜单修改时候，关联菜单err:%v", err1)
	return err
}

func (b *BaseMenu) GetBaseMenuById(id float64) (err error, menu BaseMenu) {
	err = Config.DB.Where("id = ?", id).First(&menu).Error
	return
}

func (b *BaseMenu) GetInfoList(info Services.PageInfo) (err error, list interface{}, total int) {

	err, total = Services.PagingServer(b, info)
	if err != nil {
		return
	} else {
		var menuList []BaseMenu
		err = Config.DB.Find(&menuList).Error
		return err, menuList, total
	}
}


func (m *BaseMenu) GetBaseMenuTree() (err error, menus []BaseMenu) {
	err = Config.DB.Where(" parent_id = ?", 0).Find(&menus).Error
	for i := 0; i < len(menus); i++ {
		err = getBaseChildrenList(&menus[i])
	}
	return err, menus
}

func getBaseChildrenList(menu *BaseMenu) (err error) {
	err = Config.DB.Where("parent_id = ?", menu.ID).Find(&menu.Children).Error
	for i := 0; i < len(menu.Children); i++ {
		err = getBaseChildrenList(&menu.Children[i])
	}
	return err
}
