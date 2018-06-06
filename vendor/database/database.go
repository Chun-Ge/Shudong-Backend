package database

import (
	"args"
	"crypto/md5"
	"encoding/hex"
	"entity"
	"err"

	// Register go-sql-driver for database.
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// Orm ...
var Orm *xorm.Engine

func dropConstraint() {
	// Delete Foreign Key for Table comment.
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_USER_ID")
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_POST_ID")
	Orm.Exec("alter table comment drop foreign key COMMENT_FK_NAME_LIB_ID")

	// Delete Foreign Key for Table post.
	Orm.Exec("alter table post drop foreign key POST_FK_USER_ID")
	Orm.Exec("alter table post drop foreign key POST_FK_NAME_LIB_ID")
	Orm.Exec("alter table post drop foreign key POST_FK_CATEGORY_ID")

	// Delete Foreign Key for Table name_lib.
	Orm.Exec("alter table name_lib drop foreign key NAME_LIB_FK_TOPIC_ID")

	// Delete Foreign Key for Table user_upvote_post.
	Orm.Exec("alter table user_upvote_post drop foreign key USER_UPVOTE_POST_FK_USER_ID")
	Orm.Exec("alter table user_upvote_post drop foreign key USER_UPVOTE_POST_FK_POST_ID")

	// Delete Foreign Key for Table user_upvote_comment.
	Orm.Exec("alter table user_upvote_comment drop foreign key USER_UPVOTE_COMMENT_FK_USER_ID")
	Orm.Exec("alter table user_upvote_comment drop foreign key USER_UPVOTE_COMMENT_FK_COMMENT_ID")
}

func dropTables() {
	dropConstraint()

	Orm.DropTables(&entity.AuthCode{})
	Orm.DropTables(&entity.User{})
	Orm.DropTables(&entity.Category{})
	Orm.DropTables(&entity.Comment{})
	Orm.DropTables(&entity.NameLib{})
	Orm.DropTables(&entity.Post{})
	Orm.DropTables(&entity.ReportComment{})
	Orm.DropTables(&entity.ReportPost{})
	Orm.DropTables(&entity.Topic{})
	Orm.DropTables(&entity.UserStarPost{})
	Orm.DropTables(&entity.UserUpvotePost{})
	Orm.DropTables(&entity.UserUpvoteComment{})
}

func syncTables() {
	err.CheckErr(Orm.Sync2(new(entity.AuthCode)))
	err.CheckErr(Orm.Sync2(new(entity.User)))
	err.CheckErr(Orm.Sync2(new(entity.Category)))
	err.CheckErr(Orm.Sync2(new(entity.Comment)))
	err.CheckErr(Orm.Sync2(new(entity.NameLib)))
	err.CheckErr(Orm.Sync2(new(entity.Post)))
	err.CheckErr(Orm.Sync2(new(entity.ReportComment)))
	err.CheckErr(Orm.Sync2(new(entity.ReportPost)))
	err.CheckErr(Orm.Sync2(new(entity.Topic)))
	err.CheckErr(Orm.Sync2(new(entity.UserStarPost)))
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

	// Add Foreign Key for Table report_post.
	Orm.Exec("alter table report_post add constraint REPORT_POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table report_post add constraint REPORT_POST_FK_POST_ID foreign key(post_id) REFERENCES post(id)")

	// Add Foreign Key for Table report_comment.
	Orm.Exec("alter table report_comment add constraint REPORT_COMMENT_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table report_comment add constraint REPORT_COMMENT_FK_COMMENT_ID foreign key(comment_id) REFERENCES comment(id)")

	// Add Foreign Key for Table user_star_post.
	Orm.Exec("alter table user_star_post add constraint USER_STAR_POST_FK_USER_ID foreign key(user_id) REFERENCES user(id)")
	Orm.Exec("alter table user_star_post add constraint USER_STAR_POST_FK_POST_ID foreign key(post_id) REFERENCES post(id)")

}

func encodePassword(initPassword string) (password string) {
	md5Hash := md5.New()
	md5Hash.Write([]byte(initPassword))
	password = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

func insertInitRecord() {
	Orm.Exec("truncate table user")
	Orm.Insert(&entity.User{
		Email: "1184862561@qq.com",
		// Userid:   "a2l3e4x5a6n7d8r9a0l1i",
		Password: encodePassword("123"),
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

	// init a category
	Orm.Exec("truncate table category")
	Orm.Insert(&entity.Category{
		Name: "Category-1 (init)",
	})

	// init a post
	Orm.Exec("truncate table post")
	Orm.Insert(&entity.Post{
		UserID:     1,
		CategoryID: 1,
		NameLibID:  1,
		Title:      "Title-1 (init)",
		Content:    "Post-Content-1 (init)",
		// Like: 0,
	})

	// init a comment
	Orm.Exec("truncate table comment")
	Orm.Insert(&entity.Comment{
		UserID:    1,
		PostID:    1,
		NameLibID: 1,
		Content:   "Comment-Content-1 (init)",
		// Like : 0,
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

// Start .
func Start() {
	var e error
	// strParam := "{MySQLUser}:{MySQLPassword}@tcp({MySQLURL}:{MySQLPort})/test_shudong"
	strParam := args.MySQLUser + ":" + args.MySQLPassword + "@tcp(" +
		args.MySQLURL + ":" + args.MySQLPort + ")/test_shudong"
	Orm, e = xorm.NewEngine("mysql", strParam)
	err.CheckErr(e)
	Orm.ShowSQL(false)
	Orm.SetMapper(core.GonicMapper{})

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Orm.SetDefaultCacher(cacher)

	initDatabase()
}

// StartSync .
func StartSync() {
	go Start()
}
