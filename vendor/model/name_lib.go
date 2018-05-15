package model

import (
	"database"
	"entity"
)

func GetRandomNameLib() (nameLib *entity.NameLib) {
	database.Orm.Desc("rand()").Get(&nameLib)
	return
}
