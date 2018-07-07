package model

import (
	"database"
	"testing"
)

func init() {
	database.Start()
}

func Test_GetRandomNameLib(t *testing.T) {
	name, err := GetRandomNameLib()
	if (err != nil || name == nil) {
		t.Error("Get random name lib failed.")
	} else {
		t.Log("Get random name lib passed.")
	}
}

func Test_ValidGetNameFromNameLibByID(t *testing.T) {
	var nid int64 = 2
	name := "bb"
	anotherName, err := GetNameFromNameLibByID(nid)
	if (err != nil || anotherName.Name != name || anotherName.ID != nid) {
		t.Error("Valid get name from name lib failed.")
	} else {
		t.Log("Valid get name from name lib passed.")
	}
}

// TODO: GetNameFromNameLibByID 鲁棒性
func Test_InvalidGetNameFromNameLibByID(t *testing.T) {
	var nid int64 = 7
	anotherName, err := GetNameFromNameLibByID(nid)
	if (err == nil) {
		t.Error("Invalid get name from name lib failed.")
		t.Error("Unexpected name:", anotherName)
	} else {
		t.Log("Invalid get name from name lib passed.")
	}
}