package Controllers

import (
	"github.com/gin-gonic/gin"
	"jwtapp/Models"
	"jwtapp/Services"
	"net/http"
)

type CreateAuthorityPatams struct {
	AuthorityId   string   `json:"authorityId"`
	AuthorityName string `json:"authorityName"`
}


func CreateAuthority(c *gin.Context) {
	var auth Models.Authority
	_ = c.ShouldBind(&auth)
	err, authBack := auth.CreateAuthority()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"authority": authBack,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"authority": authBack,
		})
	}
}

type DeleteAuthorityPatams struct {
	AuthorityId uint `json:"authorityId"`
}



func DeleteAuthority(c *gin.Context) {
	var a Models.Authority
	_ = c.BindJSON(&a)

	err := a.DeleteAuthority()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "delete authority success",
		})
	}
}


func GetAuthorityList(c *gin.Context){
	var pageInfo Services.PageInfo
	_ = c.BindJSON(&pageInfo)
	err, list, total := new(Models.Authority).GetInfoList(pageInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}


type GetAuthorityId struct {
	AuthorityId string `json:"authorityId"`
}


func GetAuthAndApi(c *gin.Context){
	var idInfo GetAuthorityId
	_ = c.BindJSON(&idInfo)
	err,apis := new(Models.ApiAuthority).GetAuthAndApi(idInfo.AuthorityId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"apis": apis,
		})
	}
}
