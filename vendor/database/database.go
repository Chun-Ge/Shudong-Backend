package main

import (
	"entity"
	"err"

	// Register go-sql-driver for database.
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

func dropTables() {
	orm.DropTables(&entity.User{})
	orm.DropTables(&entity.Category{})
	orm.DropTables(&entity.Comment{})
	orm.DropTables(&entity.NameLib{})
	orm.DropTables(&entity.Post{})
	orm.DropTables(&entity.Topic{})
	orm.DropTables(&entity.UserUpvotePost{})
	orm.DropTables(&entity.UserUpvoteComment{})
}

func syncTables() {
	err.CheckErr(orm.Sync2(new(entity.User)))
	err.CheckErr(orm.Sync2(new(entity.Category)))
	err.CheckErr(orm.Sync2(new(entity.Comment)))
	err.CheckErr(orm.Sync2(new(entity.NameLib)))
	err.CheckErr(orm.Sync2(new(entity.Post)))
	err.CheckErr(orm.Sync2(new(entity.Topic)))
	err.CheckErr(orm.Sync2(new(entity.UserUpvotePost)))
	err.CheckErr(orm.Sync2(new(entity.UserUpvoteComment)))
}

func main() {
	var e error

	orm, e = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/test_shudong")
	err.CheckErr(e)

	orm.SetMapper(core.GonicMapper{})

	// Clear current tables under database.
	dropTables()

	// Sync all tables.
	syncTables()
}
