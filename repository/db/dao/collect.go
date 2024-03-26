package dao

import (
	"context"
	"gorm.io/gorm"
	"meetingBooking/repository/db/model"
)

type CollectDao struct {
	*gorm.DB
}

func NewCollectDao(ctx context.Context) *CollectDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CollectDao{NewDBClient(ctx)}
}

// FindCollectProductExistOrNot 判断该用户收藏的商品是否存在
func (dao *CollectDao) FindCollectProductExistOrNot(pId, uId uint) (exists bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Collect{}).Where("product_id = ? AND user_id = ? ", pId, uId).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}
	return true, err
}
