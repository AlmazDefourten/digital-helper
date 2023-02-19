package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Application struct {
	ID int `json:"id"`

	Grant   Grant `json:"grant"`
	GrantId int   `json:"grantId"`

	Status   Status `json:"status"`
	StatusId int    `json:"statusId"`

	UserId int       `json:"userId"`
	Dates  time.Time `json:"date"`
}

// GetApplications get answer for question
func GetApplications() ([]Application, error) {
	var applications []Application
	err := db.Preload("Grant").Preload("Status").Find(&applications).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return applications, nil
}
