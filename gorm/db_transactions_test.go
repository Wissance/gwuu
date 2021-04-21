package gorm

import (
	gg "github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	g "gorm.io/gorm"
	"testing"
)

type Profile struct {
	gg.Model
	Name string `gorm:"type:varchar(128);not null;"`
}

type Role struct {
	gg.Model
	Name string `gorm:"type:varchar(128);not null;"`
}

type User struct {
	gg.Model
	UserName string `gorm:"type:varchar(128);not null;"`
	PasswordHash string `gorm:"type:varchar(255);not null;"`
	ProfileId uint
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles []Role `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //
}

type userRole struct {
	UserID  int `gorm:"primaryKey"`
	RoleID int `gorm:"primaryKey"`
}

func TestModelWithMultipleTransactions(t *testing.T) {
	cfg := g.Config{SkipDefaultTransaction: true}
	connStr := BuildConnectionString(Postgres, "127.0.0.1", 5432, "gwuu_tr_w_model_examples", dbUser, dbPassword, "disable")
	db := OpenDb2(Postgres, connStr, true, &cfg)
	assert.NotNil(t, db)

	prepareDatabase(db)
    db = db.Begin()
	// transaction 1, commit
	userProfile := Profile{Name: "user"}
	db.Create(&userProfile)

	adminProfile := Profile{Name: "admin"}
	db.Create(&adminProfile)

	regularUser := User{UserName: "def", PasswordHash: "==-b56745560", Profile: userProfile}
	db.Create(&regularUser)
	db = db.Commit()
	// transaction 2, commit

	// transaction 3, rollback

	// transaction 4 commit

    // Close
	CloseDb(db)
	// Drop
	DropDb(Postgres, connStr)
}

func prepareDatabase(db *g.DB) {
	db.AutoMigrate(Profile{})
	db.AutoMigrate(Role{})
	db.AutoMigrate(User{})
}