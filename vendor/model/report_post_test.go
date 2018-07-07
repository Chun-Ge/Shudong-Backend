package model

import (
	"database"
	"testing"

	"github.com/dmgk/faker"
)

func init() {
	database.Start()
}

func Test_ValidNewReportPost(t *testing.T) {
	var uid, pid int64 = 1, 1
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportPost(uid, pid, reason)
	if (err != nil || res != 1) {
		t.Error("Valid new report post failed.")
	} else {
		t.Log("Valid new report post passed.")
	}
}

func Test_InvalidNewReportPost_1(t *testing.T) {
	var uid, pid int64 = 1, 7
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportPost(uid, pid, reason)
	if (err == nil && res != 0) {
		t.Error("Invalid new report post #1 failed.")
	} else {
		t.Log("Invalid new report post #1 passed.")
	}
}

func Test_InvalidNewReportPost_2(t *testing.T) {
	var uid, pid int64 = 7, 1
	reason := faker.Lorem().Sentence(7)
	res, err := NewReportPost(uid, pid, reason)
	if (err == nil && res != 0) {
		t.Error("Invalid new report post #2 failed.")
	} else {
		t.Log("Invalid new report post #2 passed.")
	}
}