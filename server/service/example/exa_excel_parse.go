package example

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/xuri/excelize/v2"
)

type ExcelService struct{}

func (exa *ExcelService) ParseInfoList2Excel(infoList []autocode.JyxUser, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"报名项目", "姓名", "手机号", "证件类型", "身份证号", "性别", "民族", "省份", "城市",
		"所在单位", "出生日期", "文化程度", "考生来源", "证件领取", "职业", "年级", "班级", "参加工作时间",
		"专业年限", "政治面貌", "学历证书编号", "通讯地址", "邮寄地址", "邮箱", "户籍所在地", "原证书等级",
		"原证书编号", "原证书职业", "支付金额"})
	for i, menu := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		var birth string
		birthArr := strings.Split(menu.DateBirth, "T")
		if len(birthArr) > 0 {
			birth = birthArr[0]
		}
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			menu.ProfessionalName,
			menu.Name,
			menu.Phone,
			menu.IDType,
			menu.UID,
			menu.Sex,
			menu.Nation,
			menu.Province,
			menu.City,
			menu.CurrentUnit,
			birth,
			menu.EduLevel,
			menu.Source,
			menu.Receive,
			menu.Work,
			menu.Level,
			menu.Conditions,
			menu.WorkDate,
			menu.WorkYear,
			menu.PoliticalStatus,
			menu.SerialNumber,
			menu.Address,
			menu.PostAddress,
			menu.Email,
			menu.Place,
			menu.OCL,
			menu.OCN,
			menu.OCO,
			menu.PayAmount,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}

func (exa *ExcelService) ParseExcel2InfoList() ([]system.SysBaseMenu, error) {
	skipHeader := true
	fixedHeader := []string{"ID", "路由Name", "路由Path", "是否隐藏", "父节点", "排序", "文件名称"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "ExcelImport.xlsx")
	if err != nil {
		return nil, err
	}
	menus := make([]system.SysBaseMenu, 0)
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return nil, err
		}
		if skipHeader {
			if exa.compareStrSlice(row, fixedHeader) {
				skipHeader = false
				continue
			} else {
				return nil, errors.New("Excel格式错误")
			}
		}
		if len(row) != len(fixedHeader) {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		hidden, _ := strconv.ParseBool(row[3])
		sort, _ := strconv.Atoi(row[5])
		menu := system.SysBaseMenu{
			GVA_MODEL: global.GVA_MODEL{
				ID: uint(id),
			},
			Name:      row[1],
			Path:      row[2],
			Hidden:    hidden,
			ParentId:  row[4],
			Sort:      sort,
			Component: row[6],
		}
		menus = append(menus, menu)
	}
	return menus, nil
}

func (exa *ExcelService) compareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}
