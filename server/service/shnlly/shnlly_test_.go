package shnlly

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/shnlly"
	shnllyReq "github.com/flipped-aurora/gin-vue-admin/server/model/shnlly/request"
)

type ShnllyTestService struct {
}

// CreateShnllyTest 创建shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) CreateShnllyTest(shnllyTest *shnlly.ShnllyTest) (err error) {
	err = global.GVA_DB.Create(shnllyTest).Error
	return err
}

// DeleteShnllyTest 删除shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) DeleteShnllyTest(shnllyTest shnlly.ShnllyTest) (err error) {
	err = global.GVA_DB.Delete(&shnllyTest).Error
	return err
}

// DeleteShnllyTestByIds 批量删除shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) DeleteShnllyTestByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]shnlly.ShnllyTest{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateShnllyTest 更新shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) UpdateShnllyTest(shnllyTest shnlly.ShnllyTest) (err error) {
	err = global.GVA_DB.Save(&shnllyTest).Error
	return err
}

// GetShnllyTest 根据id获取shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) GetShnllyTest(id uint) (shnllyTest shnlly.ShnllyTest, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&shnllyTest).Error
	return
}

// GetShnllyTestInfoList 分页获取shnlly的test模型记录
// Author [piexlmax](https://github.com/piexlmax)
func (shnllyTestService *ShnllyTestService) GetShnllyTestInfoList(info shnllyReq.ShnllyTestSearch) (list []shnlly.ShnllyTest, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&shnlly.ShnllyTest{})
	var shnllyTests []shnlly.ShnllyTest
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&shnllyTests).Error
	return shnllyTests, total, err
}
