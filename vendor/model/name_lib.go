package model

import (
	"database"
	"entity"
)

// GetRandomNameLib ...
func GetRandomNameLib() (nameLib *entity.NameLib, er error) {
	_, er = database.Orm.Desc("rand()").Get(&nameLib)
	return
}

// GetNameFromNameLibByID ...
func GetNameFromNameLibByID(nameLibID int64) (ret *entity.NameLib, er error) {
	ret = &entity.NameLib{ID: nameLibID}
	_, er = database.Orm.Table("name_lib").Get(ret)
	return
}
