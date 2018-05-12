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

func addForeignKey() {
	// Add Foreign Key for Table comment.
	orm.Exec("alter table comment add constraint COMMENT_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	orm.Exec("alter table comment add constraint COMMENT_FK_POST_ID foreign key(post_id) REFERENCES post(id)")
	orm.Exec("alter table comment add constraint COMMENT_FK_NAME_LIB_ID foreign key(name_lib_id) REFERENCES name_lib(id)")

	// Add Foreign Key for Table post.
	orm.Exec("alter table post add constraint POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	orm.Exec("alter table post add constraint POST_FK_NAME_LIB_ID foreign key(name_lib_id) REFERENCES name_lib(id)")
	orm.Exec("alter table post add constraint POST_FK_CATEGORY_ID foreign key(category_id) REFERENCES category(id)")

	// Add Foreign Key for Table name_lib.
	orm.Exec("alter table name_lib add constraint NAME_LIB_FK_TOPIC_ID foreign key(topic_id) REFERENCES topic(id)")

	// Add Foreign Key for Table user_upvote_post.
	orm.Exec("alter table user_upvote_post add constraint USER_UPVOTE_POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	orm.Exec("alter table user_upvote_post add constraint USER_UPVOTE_POST_FK_POST_ID foreign key(post_id) REFERENCES post(id)")

	// Add Foreign Key for Table user_upvote_comment.
	orm.Exec("alter table user_upvote_comment add constraint USER_UPVOTE_COMMENT_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	orm.Exec("alter table user_upvote_comment add constraint USER_UPVOTE_COMMENT_FK_COMMENT_ID foreign key(comment_id) REFERENCES comment(id)")
}

func main() {
	var e error

	orm, e = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/test_shudong")
	err.CheckErr(e)
	orm.ShowSQL(true)

	orm.SetMapper(core.GonicMapper{})

	// Clear current tables under database.
	dropTables()

	// Sync all tables.
	syncTables()

	addForeignKey()
}
