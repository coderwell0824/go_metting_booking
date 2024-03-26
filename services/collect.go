package services

import (
	"context"
	"errors"
	"meetingBooking/repository/db/dao"
	"meetingBooking/utils"
)

type CollectService struct {
	ProductId   uint   `json:"productId" form:"productId"`
	BossId      uint   `json:"bossId" form:"bossId"`
	CollectId   uint   `json:"collectId" form:"collectId"`
	CollectType string `json:"collectType" form:"collectType"`
	PageNum     int    `json:"pageNum" form:"pageNum"`
	PageSize    int    `json:"pageSize" form:"pageSize"`
}

func (service *CollectService) Create(ctx context.Context) (resp interface{}, err error) {

	collectDao := dao.NewCollectDao(ctx)
	u, _ := utils.GetUserInfo(ctx)
	exists, err := collectDao.FindCollectProductExistOrNot(service.ProductId, u.Id)
	if err != nil {
		err = errors.New("查询数据库失败")
	}
	//
	if exists {
		return "该用户已收藏过该商品", nil
	}

	return
}
