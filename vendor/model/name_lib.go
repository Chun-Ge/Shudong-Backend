package model

import (
	"database"
	"entity"
	"math/rand"
)

// GetRandomNameLib ...
func GetRandomNameLib() (*entity.NameLib, error) {
	nameLib := &entity.NameLib{}

	counts, er := database.Orm.Count(nameLib)
	if er != nil {
		return nameLib, er
	}

	for true {
		randID := rand.Int63n(counts) + 1
		has, er := database.Orm.Table("name_lib").Where("id = ?", randID).Exist()
		if er != nil {
			return nameLib, er
		}
		if has {
			_, er = database.Orm.Where("id = ?", randID).Get(nameLib)
			return nameLib, er
		}
	}

	return nameLib, er
}

// GetNameFromNameLibByID ...
func GetNameFromNameLibByID(nameLibID int64) (*entity.NameLib, error) {
	ret := &entity.NameLib{ID: nameLibID}
	_, er := database.Orm.Table("name_lib").Get(ret)
	return ret, er
}
