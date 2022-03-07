package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type JyxUserService struct {
}

// CreateJyxUser 创建JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService) CreateJyxUser(jyxUser autocode.JyxUser) (err error) {
	err = global.GVA_DB.Create(&jyxUser).Error
	return err
}

// DeleteJyxUser 删除JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)DeleteJyxUser(jyxUser autocode.JyxUser) (err error) {
	err = global.GVA_DB.Delete(&jyxUser).Error
	return err
}

// DeleteJyxUserByIds 批量删除JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)DeleteJyxUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.JyxUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateJyxUser 更新JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)UpdateJyxUser(jyxUser autocode.JyxUser) (err error) {
	err = global.GVA_DB.Save(&jyxUser).Error
	return err
}

// GetJyxUser 根据id获取JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)GetJyxUser(id uint) (err error, jyxUser autocode.JyxUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&jyxUser).Error
	return
}

// GetJyxUserInfoList 分页获取JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)GetJyxUserInfoList(info autoCodeReq.JyxUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.JyxUser{})
	if info.UID != "" {
		db = db.Where("UID = ?", info.UID)
	}
	if info.ProfessionalName != "" {
		db = db.Where("ProfessionalName = ?", info.ProfessionalName)
	}
    var jyxUsers []autocode.JyxUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&jyxUsers).Error
	return err, jyxUsers, total
}
