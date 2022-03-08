package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []autocode.JyxUser `json:"infoList"`
}
