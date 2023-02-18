package test

import (
	"fmt"
	"github.com/AlmazDefourten/digital-helper/pkg/setting"
	"github.com/jinzhu/gorm"
)

func (dataTaker intf.DataTaker) TakeGrants() {

}

func FillData() {
	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Dbname))
	db.Exec("")
}
