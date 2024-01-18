// 自动生成模板ShnllyTest
package shnlly

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// shnlly的test模型 结构体  ShnllyTest
type ShnllyTest struct {
	global.GVA_MODEL
	ShnllyTestF string `json:"shnllyTestF" form:"shnllyTestF" gorm:"column:shnlly_test_f;comment:;"` //测试字段F
}

// TableName shnlly的test模型 ShnllyTest自定义表名 shnlly_test
func (ShnllyTest) TableName() string {
	return "shnlly_test"
}
