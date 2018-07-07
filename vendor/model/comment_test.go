package model

import (
	"database"
	_ "entity"
	"testing"

	"github.com/dmgk/faker"
)

func init() {
	database.Start();
}

func Test_ValidNewCommentWithRandomName(t *testing.T) {
	var uid, pid int64 = 1, 1
	content := faker.Lorem().Paragraph(7)
	_, err := NewCommentWithRandomName(uid, pid, content)
	if (err != nil) {
		t.Error("Valid new comment with random name failed.")
	} else {
		num, err := database.Orm.Table("Comment").Count()
		if (err != nil || num != 2) {
			t.Error("Valid new comment with random name failed.")
			t.Error("Current comment table item count:", num)
		} else {
			t.Log("Valid new comment with random name passed.")
		}
	}
}

// 现在是碰到error直接panic掉了，我这边err直接接不到error。。。
func Test_InvalidNewCommentWithRandomName_1(t *testing.T) {
	var uid, pid int64 = 1, 7
	content := faker.Lorem().Paragraph(7)
	_, err := NewCommentWithRandomName(uid, pid, content)
	if (err == nil) {
		t.Error("Invalid new comment with random name #1 failed.")
	} else {
		num, err := database.Orm.Table("Comment").Count()
		if (err != nil || num != 1) {
			t.Error("Invalid new comment with random name #1 failed.")
			t.Error("Current comment table item count:", num)
		} else {
			t.Log("Invalid new comment with random name #1 passed.")
		}
	}
}

func Test_InvalidNewCommentWithRandomName_2(t *testing.T) {
	var uid, pid int64 = 7, 1
	content := faker.Lorem().Paragraph(7)
	_, err := NewCommentWithRandomName(uid, pid, content)
	if (err == nil) {
		t.Error("Invalid new comment with random name #2 failed.")
	} else {
		num, err := database.Orm.Table("Comment").Count()
		if (err != nil || num != 1) {
			t.Error("Invalid new comment with random name #2 failed.")
			t.Error("Current comment table item count:", num)
		} else {
			t.Log("Invalid new comment with random name #2 passed.")
		}
	}
}

func Test_ValidCheckCommentByPost(t *testing.T) {
	var uid, cid int64 = 1, 1 
	res, err := CheckCommentByPost(uid, cid)
	if (err != nil || !res) {
		t.Error("Valid check comment by post failed.")
	} else {
		t.Log("Valid check comment by post passed.")
	}
}

func Test_InvalidCheckCommentByPost_1(t *testing.T) {
	var uid, cid int64 = 1, 7
	res, err := CheckCommentByPost(uid, cid)
	if (err == nil && res) {
		t.Error("Invalid check comment by post #1 failed.")
	} else {
		t.Log("Invalid check comment by post #1 passed.")
	}
}

func Test_InvalidCheckCommentByPost_2(t *testing.T) {
	var uid, cid int64 = 7, 1
	res, err := CheckCommentByPost(uid, cid)
	if (err == nil && res) {
		t.Error("Invalid check comment by post #2 failed.")
	} else {
		t.Log("Invalid check comment by post #2 passed.")
	}
}

func Test_ValidCheckCommentByID(t *testing.T) {
	var cid int64 = 1 
	res, err := CheckCommentByID(cid)
	if (err != nil || !res) {
		t.Error("Valid check comment by id failed.")
	} else {
		t.Log("Valid check comment by id passed.")
	}
}

func Test_InvalidCheckCommentByID(t *testing.T) {
	var cid int64 = 7
	res, err := CheckCommentByID(cid)
	if (err == nil && res) {
		t.Error("Invalid check comment by id failed.")
	} else {
		t.Log("Invalid check comment by id passed.")
	}
}

func Test_ValidCancelCommentByPost(t *testing.T) {
	var pid int64 = 1
	res, err := CancelCommentByPost(pid)
	if (err != nil || res != 1) {
		t.Error("Valid cancel comment by post failed.")
	} else {
		t.Log("Valid cancel comment by post passed.")
	}
}

func Test_InvalidCancelCommentByPost(t *testing.T) {
	var pid int64 = 7
	res, err := CancelCommentByPost(pid)
	if (err == nil && res != 0) {
		t.Error("Invalid cancel comment by post failed.")
	} else {
		t.Log("Invalid cancel comment by post passed.")
	}
}

func Test_ValidCancelCommentByID(t *testing.T) {
	var cid int64 = 1
	res, err := CancelCommentByID(cid)
	if (err != nil || res != 1) {
		t.Error("Valid cancel comment by id failed.")
	} else {
		t.Log("Valid cancel comment by id passed.")
	}
}

func Test_InvalidCancelCommentByID(t *testing.T) {
	var cid int64 = 7
	res, err := CancelCommentByID(cid)
	if (err == nil && res != 0) {
		t.Error("Invalid cancel comment by id failed.")
	} else {
		t.Log("Invalid cancel comment by id passed.")
	}
}

func Test_ValidCountCommentsOfPost(t *testing.T) {
	var pid int64 = 1
	res, err := CountCommentsOfPost(pid)
	if (err != nil || res != 1) {
		t.Error("Valid count comments of post failed.")
	} else {
		t.Log("Valid count comments of post passed.")
	}
}

func Test_InvalidCountCommentsOfPost(t *testing.T) {
	var pid int64 = 7
	res, err := CountCommentsOfPost(pid)
	if (err == nil && res != 0) {
		t.Error("Invalid count comments of post failed.")
	} else {
		t.Log("Invalid count comments of post passed.")
	}
}


func Test_ValidGetCommentsByPostID(t *testing.T) {
	var pid int64 = 1
	content := "Comment-Content-1 (init)"
	res, err := GetCommentsByPostID(pid)
	if (err != nil || len(res) != 1 || res[0].Content != content || res[0].PostID != pid) {
		t.Error("Valid get comments by post id failed.")
	} else {
		t.Log("Valid get comments by post id passed.")
	}
}

func Test_InvalidGetCommentsByPostID(t *testing.T) {
	var pid int64 = 7
	res, err := GetCommentsByPostID(pid)
	if (err == nil && len(res) != 0) {
		t.Error("Invalid count comments of post id failed.")
	} else {
		t.Log("Invalid count comments of post id passed.")
	}
}

func Test_ValidGetCommentsByUserID(t *testing.T) {
	var uid int64 = 1
	res, err := GetCommentsByUserID(uid)
	if (err != nil || len(res) != 1) {
		t.Error("Valid count comments of user id failed.")
	} else {
		t.Log("Valid count comments of user id passed.")
	}
}

func Test_InvalidGetCommentsByUserID(t *testing.T) {
	var uid int64 = 7
	res, err := GetCommentsByUserID(uid)
	if (err == nil && len(res) != 0) {
		t.Error("Invalid count comments of user id failed.")
	} else {
		t.Log("Invalid count comments of user id passed.")
	}
}