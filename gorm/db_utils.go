package gorm

import (
	"github.com/wissance/stringFormatter"
	"gorm.io/gorm"
)

const defaultPageSize = 25
//const maximumPageSize = 100

// Paginate
/* Function for getting data portion (page) by means of GORM
 * Parameters:
 *    - page - number of page starting from 1
 *    - size - number of rows to select
 * Returns gorm.DB address of database context object
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

// GetNextTableId
/* Function for getting next free id (uint) for model that have Model field with ID
 * Parameters:
 *    - db - gorm.DB address of database context object
 *    - table - table name
 * Returns MAX(ID) + 1
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