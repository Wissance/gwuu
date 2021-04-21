package gorm

import (
	"github.com/jinzhu/gorm"
	g "gorm.io/gorm"
	"testing"
)

type user struct {
	g.Model
	UserName string `gorm:"type:varchar(128);not null;"`
	PasswordHash string `gorm:"type:varchar(255);not null;"`
	ProfileId uint
	Profile profile `gorm:"foreignkey:ProfileID;"`
	Roles []roles `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type profile struct {
	g.Model
	Name string `gorm:"type:varchar(128);not null;"`
}

type roles struct {
	g.Model
	Name string `gorm:"type:varchar(128);not null;"`
}

func TestModelWithMultipleTransactions(t *testing.T) {
}

func prepareDatabase(db *g.DB) {

}