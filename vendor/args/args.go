package args

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/robfig/cron"
)

var (
	// Port ..
	Port = "8080"

	// MySQLURL .
	MySQLURL = "localhost"

	// MySQLPort .
	MySQLPort = "3306"

	// MySQLUser .
	MySQLUser = "root"

	// MySQLPassword .
	MySQLPassword = "root"

	// DEBUG ...
	DEBUG = true

	// SecretKey ...
	SecretKey = "My Secret"

	// AuthCodeSize ...
	AuthCodeSize = 6

	// AuthCodeLifeTime is the life time of auth code (minutes)
	// default to 30.
	AuthCodeLifeTime = 30

	// DeleteLogFileOnExit ...
	DeleteLogFileOnExit = false

	// JwtTokenValidDuration .
	// unit: hour
	JwtTokenValidDuration = 24
)

const (
	// TimeFormat must be specified with this time exactly.
	// "2006-01-02 15:04:05"
	// see https://www.jianshu.com/p/c7f7fbb16932
	TimeFormat = "Jan-02-2006"

	releaseLogDir = "/var/log/shudong-sysu/"
	debugLogDir   = "./log"
)

// LogDir ...
func LogDir() string {
	if DEBUG {
		return debugLogDir
	}
	return releaseLogDir
}

// UpdateVarArgs .
func UpdateVarArgs(port, mysqlURL, mysqlPort, mysqlUser, mysqlPassword string) {
	Port = port
	MySQLURL = mysqlURL
	MySQLPort = mysqlPort
	MySQLUser = mysqlUser
	MySQLPassword = mysqlPassword
}

// UpdateSecretKey .
func UpdateSecretKey() {
	c := cron.New()
	strParam := "@every " + strconv.Itoa(JwtTokenValidDuration) + "h"
	c.AddFunc(strParam, func() {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		generator := strconv.Itoa(int(rnd.Int31())) + strconv.Itoa(int(time.Now().Unix()))

		md5Hash := md5.New()
		md5Hash.Write([]byte(generator))
		SecretKey = hex.EncodeToString(md5Hash.Sum(nil))
	})
	c.Start()
}
