package main

import (
	"entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

var orm *xorm.Engine
var err error

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

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
	checkErr(orm.Sync2(new(entity.User)))
	checkErr(orm.Sync2(new(entity.Category)))
	checkErr(orm.Sync2(new(entity.Comment)))
	checkErr(orm.Sync2(new(entity.NameLib)))
	checkErr(orm.Sync2(new(entity.Post)))
	checkErr(orm.Sync2(new(entity.Topic)))
	checkErr(orm.Sync2(new(entity.UserUpvotePost)))
	checkErr(orm.Sync2(new(entity.UserUpvoteComment)))
}

func init() {
	orm, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/test_shudong")
	checkErr(err)
	orm.SetMapper(core.GonicMapper{})

	// Clear current tables under database.
	dropTables()

	// Sync all tables.
	syncTables()
}

func main() {
	app := iris.New()

	app.Run(iris.Addr(":8080"))
}
