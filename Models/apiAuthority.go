package Models


import (
	"jwtapp/Config"
	"github.com/jinzhu/gorm"

)

type ApiAuthority struct {
	gorm.Model
	AuthorityId string
	Authority   Authority `gorm:"ForeignKey:AuthorityId;AssociationForeignKey:AuthorityId"` //其实没有关联的必要
	ApiId       uint
	Api         Api
}


func (a *ApiAuthority) SetAuthAndApi(authId string, apisid []uint) (err error) {
	err = Config.DB.Where("authority_id = ?", authId).Unscoped().Delete(&ApiAuthority{}).Error
	for _, v := range apisid {
		err = Config.DB.Create(&ApiAuthority{AuthorityId: authId, ApiId: v}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取角色api关联关系
func (a *ApiAuthority) GetAuthAndApi(authId string) (err error,apiIds []uint) {
	var apis []ApiAuthority
	err = Config.DB.Where("authority_id = ?", authId).Find(&apis).Error
	for _, v := range apis {
		apiIds = append(apiIds,v.ApiId)
	}
	return nil,apiIds
}
