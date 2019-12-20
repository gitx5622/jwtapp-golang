package Models

import (
	"jwtapp/Config"
	"jwtapp/Services"
	"jwtapp/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID `json:"uuid"`
	Email    	string    `json:"email"`
	Password    string    `json:"-"`
	NickName    string    `json:"nickName" gorm:"default:'QMPlusUser'"`
	Authority   Authority  `gorm:"foreignkey:AuthorityId;association_foreignkey:id"`
	AuthorityId string    `json:"-" gorm:"default:888"`

}


func (u *User) Regist() (err error, userInter *User) {
	var user User

	findErr := Config.DB.Where("email = ?", u.Email).First(&user).Error

	if findErr == nil {
		return errors.New("Error registering"), nil
	} else {

		u.Password = utils.MD5V(u.Password)
		u.UUID = uuid.NewV4()
		u.Authority = u.Authority
		err = Config.DB.Create(u).Error

	}

	return err, u
}


//
func (u *User) ChangePassword(newPassword string) (err error, userInter *User) {
	var user User

	u.Password = utils.MD5V(u.Password)
	err = Config.DB.Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Update("password", utils.MD5V(newPassword)).Error
	return err, u
}

//
func (u *User) SetUserAuthority(uuid uuid.UUID, AuthorityId string) (err error) {
	err = Config.DB.Where("uuid = ?", uuid).First(&User{}).Update("authority_id", AuthorityId).Error
	return err
}

//
func (u *User) Login() (err error, userInter *User) {
	var user User
	u.Password = utils.MD5V(u.Password)
	err = Config.DB.Where("email = ? AND password = ?", u.Email, u.Password).First(&user).Error
	//err = Config.DB.Where("authority_id = ?", user.AuthorityId).First(&user.Authority).Error
	return err, &user
}

//
func (u *User) GetInfoList(info Services.PageInfo) (err error, list interface{}, total int) {

	err, total = Services.PagingServer(u, info)
	if err != nil {
		return
	} else {
		var userList []User
		err = Config.DB.Preload("Authority").Find(&userList).Error
		return err, userList, total
	}
}