package model

import (
	"database"
	"testing"
)

func init() {
	database.Start()
}

func Test_GetAllCategoryNames(t *testing.T) {
	// init: 已经包含一个元素
	cname := "Category-1 (init)";
	s, err := GetAllCategoryNames()
	if (err != nil || len(s) != 1 || s[0] != cname) {
		t.Error("Get all category names failed.")
	} else {
		t.Log("Get all category names passed.")
	}
}

func Test_GetCategoryNameByID(t *testing.T) {
	// init: 已经包含一个元素
	cname := "Category-1 (init)";
	c, err := GetCategoryNameByID(1)
	if (err != nil || c != cname) {
		t.Error("Get category name by id failed.")
	} else {
		t.Log("Get category name by id passed.")
	}
}

func Test_GetCategoryIDByName(t *testing.T) {
	// init: 已经包含一个元素
	cname := "Category-1 (init)";
	id, err := GetCategoryIDByName(cname)
	if (err != nil || id != 1) {
		t.Error("Get category id by name failed.")
	} else {
		t.Log("Get category id by name passed.")
	}
}