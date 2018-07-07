package database

import (
	"entity"
	"testing"

	// Register go-sql-driver for database.
	_ "github.com/go-sql-driver/mysql"

	// fake data generator
	"github.com/dmgk/faker"
)

func Test_InsertUserRecord(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.User{
		Email: faker.Internet().Email(),
		// Userid:   "a2l3e4x5a6n7d8r9a0l1i",
		Password: encodePassword(faker.Internet().Password(8, 14)),
	}); e != nil {
		t.Error("Insert user failed.")
	} else {
		t.Log("Insert user passed.")
	}
}

func Test_InsertTopicRecord_1(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.Topic{
		Name: faker.Hacker().SaySomethingSmart(),
	}); e != nil {
		t.Error("Insert topic #1 failed.")
	} else {
		t.Log("Insert topic #1 passed.")
	}
}

func Test_InsertTopicRecord_2(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.Topic{
		Name: faker.Lorem().String(),
	}); e != nil {
		t.Error("Insert topic #2 failed.")
	} else {
		t.Log("Insert topic #2 passed.")
	}
}

func Test_InsertNamelibRecord_1(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.NameLib{
		Name:    faker.Name().Name(),
		TopicID: 1,
	}); e != nil {
		t.Error("Insert namelib #1 failed.")
	} else {
		t.Log("Insert namelib #1 passed.")
	}
}

func Test_InsertNamelibRecord_2(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.NameLib{
		Name:    faker.Name().Name(),
		TopicID: 1,
	}); e != nil {
		t.Error("Insert namelib #2 failed.")
	} else {
		t.Log("Insert namelib #2 passed.")
	}
}

func Test_InsertNamelibRecord_3(t *testing.T) {
	Start()
	if _, e := Orm.Insert(&entity.NameLib{
		Name:    faker.Name().Name(),
		TopicID: 2,
	}); e != nil {
		t.Error("Insert namelib #3 failed.")
	} else {
		t.Log("Insert namelib #3 passed.")
	}
}

func Test_InsertCategory(t *testing.T) {
	Start()
	// test a category
	if _, e := Orm.Insert(&entity.Category{
		Name: "Category-1 (test)",
	}); e != nil {
		t.Error("Insert category failed.")
	} else {
		t.Log("Insert category passed.")
	}
}

func Test_InsertPost(t *testing.T) {
	Start()
	// test a post
	if _, e := Orm.Insert(&entity.Post{
		UserID:     1,
		CategoryID: 1,
		NameLibID:  1,
		Title:      faker.Lorem().String(),
		Content:    faker.Lorem().Paragraph(64),
		// Like: 0,
	}); e != nil {
		t.Error("Insert post failed.")
	} else {
		t.Log("Insert post passed.")
	}
}

func Test_InsertComment(t *testing.T) {
	Start()
	// test a comment
	if _, e := Orm.Insert(&entity.Comment{
		UserID:    1,
		PostID:    1,
		NameLibID: 1,
		Content:   faker.Lorem().Paragraph(64),
		// Like : 0,
	}); e != nil {
		t.Error("Insert comment failed.")
	} else {
		t.Log("Insert comment passed.")
	}
}