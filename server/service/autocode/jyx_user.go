package autocode

import (
	"fmt"
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

func (jyxUserService *JyxUserService)UpdatesJyxUser(jyxUser autocode.JyxUser) (err error) {
	err = global.GVA_DB.Updates(&jyxUser).Error
	return err
}

// GetJyxUser 根据id获取JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)GetJyxUser(id uint) (err error, jyxUser autocode.JyxUser) {
	err = global.GVA_DB.Where("id = ?", id).Limit(1).Find(&jyxUser).Error
	return
}

func (jyxUserService *JyxUserService)GetJyxUserByUID(uid string) (err error, jyxUser autocode.JyxUser) {
	err = global.GVA_DB.Where("uid = ?", uid).Limit(1).Find(&jyxUser).Error
	return
}

func (jyxUserService *JyxUserService)GetJyxUserByUIDAndType(uid , professionalName string) (err error, jyxUser autocode.JyxUser) {
	err = global.GVA_DB.Where("uid = ? and professionalName = ?", uid, professionalName).Limit(1).Find(&jyxUser).Error
	return
}

// GetJyxUserInfoList 分页获取JyxUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (jyxUserService *JyxUserService)GetJyxUserInfoList(info autoCodeReq.JyxUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	fmt.Println(info)
    // 创建db
	db := global.GVA_DB.Model(&autocode.JyxUser{})
	if info.UID != "" {
		db = db.Where("UID = ?", info.UID)
	}
	if info.ProfessionalName != "" {
		db = db.Where("professionalName = ?", info.ProfessionalName)
	}
	if *info.Verify > 0 {
		db = db.Where("verify = ?", *info.Verify)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
    var jyxUsers []autocode.JyxUser
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&jyxUsers).Error
	return err, jyxUsers, total
}
