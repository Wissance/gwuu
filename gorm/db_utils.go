package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/wissance/stringFormatter"
)

const defaultPageSize = 25
//const maximumPageSize = 100

/*
 * Function for getting data portion (page) by means of GORM
 */
func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if size < 0 {
			size = defaultPageSize
		}
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

/*
 * Function for getting next uint id (for model that have Model field with ID as uint)
 */
func GetNextTableId(db *gorm.DB, table string) uint {
	type Internal struct {
		Id uint
	}
	var maxId Internal
	getMaxIdQuery := stringFormatter.Format("SELECT MAX(id) As Id FROM {0};", table);
	db.Raw(getMaxIdQuery).Scan(&maxId)
	return maxId.Id + 1
}