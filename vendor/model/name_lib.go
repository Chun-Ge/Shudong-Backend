package model

import (
	"database"
	"entity"
)

// GetRandomNameLib .
func GetRandomNameLib() (nameLib *entity.NameLib) {
	database.Orm.Desc("rand()").Get(&nameLib)
	return
}
