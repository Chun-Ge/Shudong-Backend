package model

import (
	"database"
	"testing"
)

// TODO: Start之后插入初始元素没用
func init() {
	database.Start()
}

func Test_ValidNewSharePost(t *testing.T) {
	var uid, pid int64 = 1, 1
	res, err := NewSharePost(uid, pid)
	if (err != nil || res != 1) {
		t.Error("Valid new share post failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 1)
	} else if (res != 1) {
		t.Log("Valid new share post passed.")
	}
}

func Test_InvalidNewSharePost_1(t *testing.T) {
	var uid, pid int64 = 1, 7
	res, err := NewSharePost(uid, pid)
	if (err == nil && res != 0) {
		t.Error("Invalid new share post #1 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 0)
	} else {
		t.Log("Invalid new share post #1 passed.")
	}
}

func Test_InvalidNewSharePost_2(t *testing.T) {
	var uid, pid int64 = 7, 1
	res, err := NewSharePost(uid, pid)
	if (err == nil && res != 0) {
		t.Error("Invalid new share post #2 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 0)
	} else {
		t.Log("Invalid new share post #2 passed.")
	}
}

func Test_ValidCancelSharePost(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	res, err := CancelSharePost(uid, pid)
	if (err != nil || res != 1) {
		t.Error("Valid cancel share post failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 1)
	} else {
		t.Log("Valid cancel share post passed.")
	}
}

func Test_InvalidCancelSharePost_1(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	pid = 7
	res, err := CancelSharePost(uid, pid)
	if (err == nil && res != 0) {
		t.Error("Invalid cancel share post #1 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 0)
	} else {
		t.Log("Invalid cancel share post #1 passed.")
	}
}


func Test_InvalidCancelSharePost_2(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	uid = 7
	res, err := CancelSharePost(uid, pid)
	if (err == nil && res != 0) {
		t.Error("Invalid cancel share post #2 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", 0)
	} else {
		t.Log("Invalid cancel share post #2 passed.")
	}
}

func Test_CountPostShared(t *testing.T) {
	var uid, pid int64 = 1, 1
	var round, i int64 = 3, 1
	for i = 1; i <= round; i++ {
		NewSharePost(uid, pid)
	}
	res, err := CountPostShared(pid)
	if (err != nil || res != round) {
		t.Error("Count post shared failed.")
		t.Error("Result:", res)
		t.Error("Expect:", round)
	} else {
		t.Log("Count post shared passed.")
	}
}

func Test_ValidCheckPostIfShared(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	res, err := CheckPostIfShared(uid, pid)
	if (err != nil || !res) {
		t.Error("Valid check post if shared failed.")
		t.Error("Result:", res)
		t.Error("Expect:", true)
	} else {
		t.Log("Valid check post if shared passed.")
	}
}

func Test_InvalidCheckPostIfShared_1(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	pid = 7
	res, err := CheckPostIfShared(uid, pid)
	if (err == nil && res) {
		t.Error("Invalid check post if shared #1 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", false)
	} else {
		t.Log("Invalid check post if shared #1 passed.")
	}
}

func Test_InvalidCheckPostIfShared_2(t *testing.T) {
	var uid, pid int64 = 1, 1
	NewSharePost(uid, pid)
	uid = 7
	res, err := CheckPostIfShared(uid, pid)
	if (err == nil && res) {
		t.Error("Invalid check post if shared #2 failed.")
		t.Error("Result:", res)
		t.Error("Expect:", false)
	} else {
		t.Log("Invalid check post if shared #2 passed.")
	}
}