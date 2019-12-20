package Services

import (
	"jwtapp/Config"
)

type PageInfo struct {
	Page     int
	PageSize int
}

type Paging interface {
	GetInfoList(PageInfo) (err error, list interface{}, total int)
}

func PagingServer(paging Paging, info PageInfo) (err error, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err = Config.DB.Model(paging).Count(&total).Error
	Config.DB.Limit(limit).Offset(offset).Order("id desc")
	return err, total
}

