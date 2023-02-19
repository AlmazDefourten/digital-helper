package data

import (
	"fmt"
	"github.com/AlmazDefourten/digital-helper/models"
	"github.com/AlmazDefourten/digital-helper/models/intf"
	"github.com/AlmazDefourten/digital-helper/pkg/setting"
	"github.com/AlmazDefourten/digital-helper/test"
	"github.com/jinzhu/gorm"
	"log"
)

func GetGrants() ([]models.Grant, error) {
	var err error
	var db = &gorm.DB{}
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Dbname))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	var dataSource intf.DataTaker = &test.DataTest{}
	grants, err := dataSource.TakeGrants(db)
	if err != nil {
		return nil, err
	}
	return grants, nil
}
