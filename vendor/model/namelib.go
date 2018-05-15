package model

import (
	_ "database"
	"entity"
)

func GetRandomNameLib() (name *entity.NameLib, err error) {
    name = &entity.NameLib{}
    return
}
