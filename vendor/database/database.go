package database

import (
	"entity"
	"err"

	// Register go-sql-driver for database.
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// Orm .
var Orm *xorm.Engine

func dropConstraint() {
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_USER_ID")
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_POST_ID")
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_NAME_LIB_ID")

	Orm.Exec("alter table name_lib drop foreign key NAME_LIB_FK_TOPIC_ID")

	Orm.Exec("alter table user_upvote_post drop foreign key USER_UPVOTE_POST_FK_USER_ID")
	Orm.Exec("alter table user_upvote_post drop foreign key USER_UPVOTE_POST_FK_POST_ID")

	Orm.Exec("alter table user_upvote_comment drop foreign key USER_UPVOTE_COMMENT_FK_USER_ID")
	Orm.Exec("alter table user_upvote_comment drop foreign key USER_UPVOTE_COMMENT_FK_COMMENT_ID")
}

func dropTables() {
	dropConstraint()

	Orm.DropTables(&entity.User{})
	Orm.DropTables(&entity.Category{})
	Orm.DropTables(&entity.Comment{})
	Orm.DropTables(&entity.NameLib{})
	Orm.DropTables(&entity.Post{})
	Orm.DropTables(&entity.Topic{})
	Orm.DropTables(&entity.UserUpvotePost{})
	Orm.DropTables(&entity.UserUpvoteComment{})
}

func syncTables() {
	err.CheckErr(Orm.Sync2(new(entity.User)))
	err.CheckErr(Orm.Sync2(new(entity.Category)))
	err.CheckErr(Orm.Sync2(new(entity.Comment)))
	err.CheckErr(Orm.Sync2(new(entity.NameLib)))
	err.CheckErr(Orm.Sync2(new(entity.Post)))
	err.CheckErr(Orm.Sync2(new(entity.Topic)))
	err.CheckErr(Orm.Sync2(new(entity.UserUpvotePost)))
	err.CheckErr(Orm.Sync2(new(entity.UserUpvoteComment)))
}

func addForeignKey() {
	// Add Foreign Key for Table comment.
	Orm.Exec("alter table comment add constraint COMMENT_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table comment add constraint COMMENT_FK_POST_ID foreign key(post_id) REFERENCES post(id)")
	Orm.Exec("alter table comment add constraint COMMENT_FK_NAME_LIB_ID foreign key(name_lib_id) REFERENCES name_lib(id)")

	// Add Foreign Key for Table post.
	Orm.Exec("alter table post add constraint POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table post add constraint POST_FK_NAME_LIB_ID foreign key(name_lib_id) REFERENCES name_lib(id)")
	Orm.Exec("alter table post add constraint POST_FK_CATEGORY_ID foreign key(category_id) REFERENCES category(id)")

	// Add Foreign Key for Table name_lib.
	Orm.Exec("alter table name_lib add constraint NAME_LIB_FK_TOPIC_ID foreign key(topic_id) REFERENCES topic(id)")

	// Add Foreign Key for Table user_upvote_post.
	Orm.Exec("alter table user_upvote_post add constraint USER_UPVOTE_POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table user_upvote_post add constraint USER_UPVOTE_POST_FK_POST_ID foreign key(post_id) REFERENCES post(id)")

	// Add Foreign Key for Table user_upvote_comment.
	Orm.Exec("alter table user_upvote_comment add constraint USER_UPVOTE_COMMENT_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table user_upvote_comment add constraint USER_UPVOTE_COMMENT_FK_COMMENT_ID foreign key(comment_id) REFERENCES comment(id)")
}

func insertInitRecord() {
	Orm.Exec("truncate table user")
	Orm.Insert(&entity.User{
		Email: "1184862561@qq.com",
		// Userid:   "a2l3e4x5a6n7d8r9a0l1i",
		Password: "123",
	})

	Orm.Exec("truncate table topic")
	Orm.Insert(&entity.Topic{
		Name: "AA",
	})
	Orm.Insert(&entity.Topic{
		Name: "BB",
	})

	Orm.Exec("truncate table name_lib")
	Orm.Insert(&entity.NameLib{
		Name:    "aa",
		TopicID: 1,
	})
	Orm.Insert(&entity.NameLib{
		Name:    "bb",
		TopicID: 1,
	})

	Orm.Insert(&entity.NameLib{
		Name:    "cc",
		TopicID: 2,
	})

}

func initDatabase() {
	// Clear current tables under database.
	dropTables()
	// Sync all tables.
	syncTables()
	insertInitRecord()
	addForeignKey()
}

func init() {
	var e error

	Orm, e = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/test_shudong")
	err.CheckErr(e)
	Orm.ShowSQL(false)
	Orm.SetMapper(core.GonicMapper{})

	initDatabase()
}
