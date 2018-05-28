package args

var (
	// Port ..
	Port = ":8080"
)

const (
	// SecretKey ..
	SecretKey = "My Secret"
	// this const arg should in args/args.go

	// auth code size
	Auth_Code_Size = 6
	// the lefe-time of auth code, default: 30minutes
	Auth_Code_Life_Time = 30
)
