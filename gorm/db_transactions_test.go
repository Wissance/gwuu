package gorm

import (
	"errors"
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
	UserName     string `gorm:"type:varchar(128);not null;"`
	PasswordHash string `gorm:"type:varchar(255);not null;"`
	ProfileId    uint
	Profile      Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Roles        []Role  `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` //
}

func TestModelWithMultipleNestedTransactions(t *testing.T) {
	cfg := g.Config{SkipDefaultTransaction: true}
	connStr := BuildConnectionString(Postgres, "127.0.0.1", 5432, "gwuu_tr_w_model_examples", dbUser, dbPassword, "disable")
	db := OpenDb2(Postgres, connStr, true, true, &cfg, &postgresCollation)
	assert.NotNil(t, db)

	prepareDatabase(db)
	var regularUser User
	db.Transaction(func(tx *g.DB) error {
		// transaction 1, commit
		userProfile := Profile{Name: "user"}
		db.Create(&userProfile)

		adminProfile := Profile{Name: "admin"}
		db.Create(&adminProfile)

		regularUser = User{UserName: "def", PasswordHash: "==-b56745560", Profile: userProfile}
		db.Create(&regularUser)
		return nil
	})
	var user User
	db.Model(User{}).Where("user_name = ?", regularUser.UserName).First(&user)
	assert.True(t, user.ID > 0)
	// transaction 2, commit
	db.Transaction(func(tx *g.DB) error {
		role1 := Role{Name: "regular"}
		db.Create(&role1)
		role2 := Role{Name: "administrator"}
		db.Create(&role2)
		return nil
	})
	var role Role
	db.Model(Role{}).Where("name = ?", "regular").First(&role)
	assert.True(t, role.ID > 0)
	db.Model(Role{}).Where("name = ?", "administrator").First(&role)
	assert.True(t, role.ID > 0)
	// transaction 3, rollback
	db.Transaction(func(tx *g.DB) error {
		db.Delete(&role)
		return errors.New("just testing rollback")
	})
	db.Model(Role{}).Where("name = ?", "administrator").First(&role)
	assert.True(t, role.ID > 0)
	// Close
	CloseDb(db)
	// Drop
	DropDb(Postgres, connStr, &cfg)
}

func TestModelWithMultipleManualTransactions(t *testing.T) {
	cfg := g.Config{SkipDefaultTransaction: true}
	connStr := BuildConnectionString(Postgres, "127.0.0.1", 5432, "gwuu_tr_w_model_examples", dbUser, dbPassword, "disable")
	db := OpenDb2(Postgres, connStr, true, true, &cfg, &postgresCollation)
	assert.NotNil(t, db)

	prepareDatabase(db)
	sessionConfig := g.Session{SkipDefaultTransaction: true}
	tx := db.Session(&sessionConfig)
	tx.Begin()
	role1 := Role{Name: "regular"}
	tx.Create(&role1)
	role2 := Role{Name: "administrator"}
	tx.Create(&role2)
	tx.Commit()
	var role Role
	db.Model(Role{}).Where("name = ?", role1.Name).First(&role)
	assert.True(t, role.ID > 0)
	db.Model(Role{}).Where("name = ?", role2.Name).First(&role)
	assert.True(t, role.ID > 0)
	tx.Begin()
	userProfile := Profile{Name: "user"}
	db.Create(&userProfile)
	adminProfile := Profile{Name: "admin"}
	db.Create(&adminProfile)
	tx.Commit()
	var profile Profile
	db.Model(Profile{}).Where("name = ?", userProfile.Name).First(&profile)
	assert.True(t, role.ID > 0)
	db.Model(Profile{}).Where("name = ?", adminProfile.Name).First(&profile)
	assert.True(t, role.ID > 0)
	// Close
	CloseDb(db)
	// Drop
	DropDb(Postgres, connStr, &cfg)
}

func prepareDatabase(db *g.DB) {
	db.AutoMigrate(Profile{})
	db.AutoMigrate(Role{})
	db.AutoMigrate(User{})
}
