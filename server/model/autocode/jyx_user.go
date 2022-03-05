// 自动生成模板JyxUser
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// JyxUser 结构体
// 如果含有time.Time 请自行import time包
type JyxUser struct {
      global.GVA_MODEL
      Address  string `json:"address" form:"address" gorm:"column:address;comment:通讯地址;size:256;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:城市;size:64;"`
      Conditions  string `json:"conditions" form:"conditions" gorm:"column:conditions;comment:申报条件;size:64;"`
      CurrentUnit  string `json:"currentUnit" form:"currentUnit" gorm:"column:currentUnit;comment:所在单位;size:64;"`
      DateBirth  string `json:"dateBirth" form:"dateBirth" gorm:"column:dateBirth;comment:出生日期;size:64;"`
      EduLevel  string `json:"eduLevel" form:"eduLevel" gorm:"column:eduLevel;comment:文化程度;size:64;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:64;"`
      Experience  string `json:"experience" form:"experience" gorm:"column:experience;comment:简要经历;"`
      IDType  string `json:"IDType" form:"IDType" gorm:"column:IDType;comment:证件类型;size:64;"`
      Level  string `json:"level" form:"level" gorm:"column:level;comment:级别;size:64;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:姓名;size:64;"`
      Nation  string `json:"nation" form:"nation" gorm:"column:nation;comment:民族;size:64;"`
      OCL  string `json:"OCL" form:"OCL" gorm:"column:OCL;comment:原证书等级;size:64;"`
      OCN  string `json:"OCN" form:"OCN" gorm:"column:OCN;comment:原证书编号;size:64;"`
      OCO  string `json:"OCO" form:"OCO" gorm:"column:OCO;comment:原证书职业;size:64;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:64;"`
      Place  string `json:"place" form:"place" gorm:"column:place;comment:户籍所在地;size:256;"`
      PoliticalStatus  string `json:"politicalStatus" form:"politicalStatus" gorm:"column:politicalStatus;comment:政治面貌;size:64;"`
      PostAddress  string `json:"postAddress" form:"postAddress" gorm:"column:postAddress;comment:邮寄地址;size:256;"`
      PostalCode  string `json:"postalCode" form:"postalCode" gorm:"column:postalCode;comment:邮政编码;size:64;"`
      ProfessionalName  string `json:"professionalName" form:"professionalName" gorm:"column:professionalName;comment:职业名称;size:64;"`
      Province  string `json:"province" form:"province" gorm:"column:province;comment:省份;size:64;"`
      Receive  string `json:"receive" form:"receive" gorm:"column:receive;comment:证件领取;size:64;"`
      SerialNumber  string `json:"serialNumber" form:"serialNumber" gorm:"column:serialNumber;comment:学历证书编号;size:64;"`
      Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:64;"`
      Source  string `json:"source" form:"source" gorm:"column:source;comment:考生来源;size:64;"`
      UID  string `json:"UID" form:"UID" gorm:"column:UID;comment:身份证号;size:64;"`
      Work  string `json:"work" form:"work" gorm:"column:work;comment:从事专业;size:64;"`
      WorkDate  string `json:"workDate" form:"workDate" gorm:"column:workDate;comment:参加工作时间;size:64;"`
      WorkType  string `json:"workType" form:"workType" gorm:"column:workType;comment:工种名称;size:64;"`
      WorkYear  string `json:"workYear" form:"workYear" gorm:"column:workYear;comment:专业年限;size:64;"`
      UserCertify  string `json:"userCertify" form:"userCertify" gorm:"column:userCertify;comment:学籍在线证明或工作证明;size:600;"`
      UserPic  string `json:"userPic" form:"userPic" gorm:"column:userPic;comment:用户照片;size:600;"`
      UserPay  string `json:"userPay" form:"userPay" gorm:"column:userPay;comment:支付凭证;size:600;"`
}


// TableName JyxUser 表名
func (JyxUser) TableName() string {
  return "jyx_user"
}

