package model

import (
	"database"
	"entity"
	"testing"

	"github.com/dmgk/faker"
)

func authcodeEqual(a *entity.AuthCode, b *entity.AuthCode) bool {
	return a.ID == b.ID && a.UserID == b.UserID && a.Code == b.Code;
}

func init() {
	database.Start()
}

func Test_ValidAuthCode(t *testing.T) {
	var uid, acid int64 = 1, 1
	code := faker.Number().Number(6)
	ac, e := NewAuthCode(uid, code)
	if e != nil {
		t.Error("New auth code failed.")
		t.FailNow()
	} else {
		t.Log("New auth code passed.")
	}
	// test CheckAuthCodeByUser #1
	if ok, e := CheckAuthCodeByUser(uid); e != nil || !ok {
		t.Error("Check auth code by user #1 failed.")
	} else {
		t.Log("Check auth code by user #1 passed.")
	}
	// test CheckAuthCodeByUser #2
	var anotheruid int64
	for anotheruid = -3; anotheruid <= 3; anotheruid++ {
		if anotheruid == uid {
			continue
		}
		if ok, e := CheckAuthCodeByUser(anotheruid); e == nil && ok {
			t.Error("Check auth code by user #2 failed: uid =",anotheruid,".")
		} else {
			t.Log("Check auth code by user #2 passed: uid =",anotheruid,".")
		}
	}
	// test CheckAuthCodeByUserAndCode #1
	if ok, e := CheckAuthCodeByUserAndCode(uid, code); e != nil || !ok {
		t.Error("Check auth code by user #1 failed.")
	} else {
		t.Log("Check auth code by user #1 passed.")
	}
	// test CheckAuthCodeByUserAndCode #2
	for {
		anotherCode := faker.Number().Number(6)
		if (anotherCode != code) {
			if ok, e := CheckAuthCodeByUserAndCode(uid, anotherCode); e == nil && ok {
				t.Error("Check auth code by user #2 failed.")
			} else {
				t.Log("Check auth code by user #2 passed.")
			}
			break
		}
	}
	// test CheckAuthCodeByUserAndCode #3
	for anotheruid = -3; anotheruid <= 3; anotheruid++ {
		if anotheruid == uid {
			continue
		}
		if ok, e := CheckAuthCodeByUserAndCode(anotheruid, code); e == nil && ok {
			t.Error("Check auth code by user #3 failed: uid =",anotheruid,".")
		} else {
			t.Log("Check auth code by user #3 passed: uid =",anotheruid,".")
		}
	}
	// test GetAuthCodeByUserAndCode #1
	if acValidate, ok, e := GetAuthCodeByUserAndCode(uid, code); e != nil || !ok || !authcodeEqual(ac,acValidate) {
		t.Error("Get auth code by user #1 failed.")
		t.Error("Result:", acValidate)
		t.Error("Expect:", ac)
	} else {
		t.Log("Get auth code by user #1 passed.")
	}
	// test GetAuthCodeByUserAndCode #2
	for {
		anotherCode := faker.Number().Number(6)
		if (anotherCode != code) {
			if acValidate, ok, e := GetAuthCodeByUserAndCode(uid, anotherCode); e == nil && ok {
				t.Error("Get auth code by user #2 failed.")
				t.Error("Result:", acValidate)
				t.Error("Expect: None")
			} else {
				t.Log("Get auth code by user #2 passed.")
			}
			break
		}
	}
	// test GetAuthCodeByUserAndCode #3
	for anotheruid = -3; anotheruid <= 3; anotheruid++ {
		if anotheruid == uid {
			continue
		}
		if acValidate, ok, e := GetAuthCodeByUserAndCode(anotheruid, code); e == nil && ok {
			t.Error("Get auth code by user #3 failed: uid =",anotheruid,".")
			t.Error("Result:", acValidate)
			t.Error("Expect: None")
		} else {
			t.Log("Get auth code by user #3 passed: uid =",anotheruid,".")
		}
	}
	// test UpdateAuthCode #1
	for {
		anotherCode := faker.Number().Number(6)
		if (anotherCode != code) {
			// 更新是否可行
			if e := UpdateAuthCode(uid, anotherCode); e != nil {
				t.Error("Update auth code #1 failed.")
			} else {
				t.Log("Update auth code #1 passed.")
			}
			// 更新是否查得到
			if ok, e := CheckAuthCodeByUserAndCode(uid, anotherCode); e != nil || !ok {
				t.Error("Auth code update #1 check failed.")
			} else {
				t.Log("Auth code update #1 check passed.")
			}
			break
		}
	}
	// test UpdateAuthCode #2
	for anotheruid = -3; anotheruid <= 3; anotheruid++ {
		if anotheruid == uid {
			continue
		}
		for {
			anotherCode := faker.Number().Number(6)
			if (anotherCode != code) {
				// 更新是否可行
				// 对于不存在的uid，UpdateAuthCode不会有数据库的修改
				if e := UpdateAuthCode(anotheruid, anotherCode); e != nil {
					t.Error("Update auth code #2 failed: uid =",anotheruid,".")
				} else {
					t.Log("Update auth code #2 passed: uid =",anotheruid,".")
				}
				// 更新是否查得到（应该是查不到才对）
				if ok, e := CheckAuthCodeByUserAndCode(anotheruid, anotherCode); e == nil && ok {
					t.Error("Auth code update #2 check failed: uid =",anotheruid,".")
				} else {
					t.Log("Auth code update #2 check passed: uid =",anotheruid,".")
				}
				break
			}
		}
	}
	// test DeleteAuthCode #1
	if affected, e := DeleteAuthCode(acid); e != nil || affected != 1 {
		t.Error("Delete auth code #1 failed.")
	} else {
		t.Log("Delete auth code #1 passed.")
	}
	// test DeleteAuthCode #2
	if affected, e := DeleteAuthCode(acid); e != nil || affected != 0 {
		t.Error("Delete auth code #2 failed.")
	} else {
		t.Log("Delete auth code #2 passed.")
	}
	// test DeleteAuthCode #3
	if affected, e := DeleteAuthCode(-1); e != nil || affected != 0 {
		t.Error("Delete auth code #3 failed.")
	} else {
		t.Log("Delete auth code #3 passed.")
	}
}