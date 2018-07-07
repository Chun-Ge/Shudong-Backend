package model

import (
	"database"
	"testing"

	"github.com/dmgk/faker"
)

func init() {
	database.Start()
}

func Test_ValidNewReportComment(t *testing.T) {
	var uid, cid int64 = 1, 1
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportComment(uid, cid, reason)
	if (err != nil || res != 1) {
		t.Error("Valid new report comment failed.")
	} else {
		t.Log("Valid new report comment passed.")
	}
}

func Test_InvalidNewReportComment_1(t *testing.T) {
	var uid, cid int64 = 1, 7
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportPost(uid, cid, reason)
	if (err == nil && res != 0) {
		t.Error("Invalid new report comment #1 failed.")
	} else {
		t.Log("Invalid new report comment #1 passed.")
	}
}

func Test_InvalidNewReportComment_2(t *testing.T) {
	var uid, cid int64 = 7, 1
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportPost(uid, cid, reason)
	if (err == nil && res != 0) {
		t.Error("Invalid new report comment #2 failed.")
	} else {
		t.Log("Invalid new report comment #2 passed.")
	}
}