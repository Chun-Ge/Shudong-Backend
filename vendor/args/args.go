package args

var (
	// Port ..
	Port = ":8080"
)

const (
	// SecretKey ..
	SecretKey = "My Secret"

	// DEBUG ..
	DEBUG = true

	// DeleteLogFileOnExit ..
	DeleteLogFileOnExit = false

	releaseLogDir = "/var/log/shudong-sysu/"
)

// LogDir .
func LogDir() string {
	if DEBUG {
		return "./"
	}
	return releaseLogDir
}
