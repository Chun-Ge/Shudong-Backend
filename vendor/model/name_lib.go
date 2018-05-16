package model

import (
	"database"
	"entity"
)

// GetRandomNameLib .
func GetRandomNameLib() (nameLib *entity.NameLib, er error) {
	_, er = database.Orm.Desc("rand()").Get(&nameLib)
	return
}
