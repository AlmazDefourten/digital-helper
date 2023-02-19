package intf

import (
	"github.com/AlmazDefourten/digital-helper/models"
	"github.com/jinzhu/gorm"
)

type DataTaker interface {
	TakeGrants(db *gorm.DB) ([]models.Grant, error)
}
