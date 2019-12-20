package Controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"jwtapp/Middlewares"
	"jwtapp/Models"
	"jwtapp/Services"
	"net/http"
	"time"
)


type RegistAndLoginStuct struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type ChangePasswordStutrc struct {
	Email    string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}



func Regist(c *gin.Context) {
	var R RegistAndLoginStuct
	_ = c.BindJSON(&R)

	U := &Models.User{Email: R.Email, Password: R.Password}
	err, user := U.Regist()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg" : "Registration failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg" : "Registration success",
			"user": user,
		})
	}
}

func Login(c *gin.Context) {
	var L RegistAndLoginStuct
	_ = c.BindJSON(&L)
	U := &Models.User{Email: L.Email, Password: L.Password}
	if err, user := U.Login(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":"403",
			"msg":"Login unsuccessful",

		})
	} else {
		tokenNext(c, *user)
	}
}


func tokenNext(c *gin.Context, user Models.User) {
	j := &Middlewares.JWT{
		[]byte("qmPlus"),
	}

	expirationTime := time.Now().Add(5 * time.Hour)

	clams := Middlewares.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,

		StandardClaims: jwt.StandardClaims{
			//NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "qmPlus",
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg":"Token creation unsuccessful"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":"200",
			"token": token,
			"expiresAt": clams.StandardClaims.ExpiresAt * 1,
		})
	}
}



func ChangePassword(c *gin.Context) {
	var params ChangePasswordStutrc
	_ = c.BindJSON(&params)
	U := &Models.User{Email: params.Email, Password: params.Password}
	if err, _ := U.ChangePassword(params.NewPassword); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code" : "403",
			"msg":"Change Password unsuccessful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":"Change password successful",
		})
	}
}

func GetUserList(c *gin.Context) {
	var pageInfo Services.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(Models.User).GetInfoList(pageInfo)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg":"Failed to get user list"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"userList": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

func Logout(c *gin.Context)  {
	c.Request.Header.Set("","")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Logged out",
		"data": "success",
	})

}


func SetUserAuthority(c *gin.Context) {
	var sua SetUserAuth
	_ = c.BindJSON(&sua)
	err := new(Models.User).SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg":"Failed to set User security"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":"success",
		})
	}
}
