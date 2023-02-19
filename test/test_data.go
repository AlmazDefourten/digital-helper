package test

import (
	"github.com/AlmazDefourten/digital-helper/models"
	"github.com/jinzhu/gorm"
)

type DataTest struct {
}

func (testData *DataTest) TakeGrants(db *gorm.DB) ([]models.Grant, error) {
	var grants []models.Grant
	err := db.Find(&grants).Error
	if err != nil {
		return nil, err
	}

	return grants, nil
}

/*func FillData() {
	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Dbname))
	if err != nil {
		log.Println(err)
	}
}*/
